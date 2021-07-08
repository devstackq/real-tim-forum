import Parent from "./Parent.js";

export default class Chat extends Parent {
  constructor() {
    super();
    this.ws = new WebSocket("ws://localhost:6969/api/chat");
    this.users = [];
    this.chatbox = document.getElementById("chatbox");
  }

  setTitle(title) {
    document.title = title;
  }
  //onscroll get offset messages
  render(seq, where) {
    let parent = document.querySelector(where);
    seq.forEach((item) => {
      let div = document.createElement("div");
      for (let [i, v] of Object.entries(item)) {
        let span = document.createElement("span");
        span.textContent = ` ${i} : ${v} `;
        div.append(span);
      }
      parent.append(div);
    });
  }
  //uuid receiver & send my uuid
  //show all user - except yourself todo:

  showOnlineUsers(object) {
    let parent = document.getElementById("userlistbox");
    let ul = document.createElement("ul");
    parent.innerHTML = "";
    ul.id = "listusersID";
    // iter obj users
    for (let [k, user] of Object.entries(object)) {
      this.users.push(user);
      let li = "";
      for (let [key, value] of Object.entries(user)) {
        if (Object.entries(user).length == 1) {
          super.showNotify("Now, no has online user", "error");
          return;
        }
        if (key == "fullname") {
          li = document.createElement("li");
          li.textContent = value;

          li.onclick = (e) => {
            // chatbox.innerHTML = "";
            let obj = {
              receiver: user["UUID"],
              sender: super.getUserSession(),
              type: "getmessages",
            };
            this.ws.send(JSON.stringify(obj));
          };
        }

        if (
          user["UUID"] != super.getUserSession() &&
          Object.entries(user).length > 1
        ) {
          ul.append(li);
        }
      }
    }
    parent.append(ul);
  }

  sendMessage(receiver) {
    let chatbox = document.querySelector("#chatbox");
    let uid = super.getUserId();
    let content = document.getElementById("messageFieldId").value;
    let senderUUID = super.getUserSession();
    let message = {
      content: content,
      sender: senderUUID,
      receiver: receiver,
      userid: parseInt(uid),
      type: "newmessage",
    };
    let senderName = "";
    for (let [k, v] of Object.entries(this.users)) {
      if (k === senderUUID) {
        senderName = v;
      }
    }
    let div = document.createElement("div");
    let span = document.createElement("span");
    span.className = "chat_sender";
    span.textContent = ` ${message.content} ${new Date()} ${senderName} `;
    // chatbox.children[chatbox.children.length - 3].append(span);
    div.append(span);
    if (chatbox.children != undefined) {
      if (chatbox.children.length > 2) {
        chatbox.children[chatbox.children.length - 3].append(div);
      } else {
        chatbox.children[chatbox.children.length - 2].append(div);
      }
    }
    document.getElementById("messageFieldId").value = "";

    this.ws.send(JSON.stringify(message));
    // chatbox.children[chatbox.children.length - 2].append(li);
  }
  //DRY ?
  showListMessage(messages) {
    let userid = super.getUserId();
    if (messages != null) {
      let chat = document.querySelector("#chatbox");
      chat.innerHTML = "";

      messages.forEach((item) => {
        let div = document.createElement("div");
        for (let [k, v] of Object.entries(item)) {
          let span = document.createElement("span");
          if (v != null && v != "" && k != "sender" && k != "receiver") {
            span.textContent = `  ${v}: `;
          }
          if (k == "userid") {
            if (v == userid) {
              div.classList.add("chat_sender");
            }
          }
          div.append(span);
        }
        chat.append(div);
      });
      //call func
      let receive = "";

      if (messages.length != 0) {
        receive = messages[0]["receiver"];
      }
      this.showChatWindow(this, receive);
    }
  }
  //dr, 00, msg.receivery
  showChatWindow(scope, receiver) {
    document.getElementById("chatbox").style.display = "block";
    //dry?
    let chat = document.querySelector("#chatbox");
    let textarea = document.createElement("textarea");
    let sendBtn = document.createElement("button");

    textarea.id = "messageFieldId";
    sendBtn.id = "sendBtnId";
    sendBtn.textContent = "send messageq";
    // console.log(receiver, 11, this)
    sendBtn.onclick = this.sendMessage.bind(this, receiver);
    chat.append(textarea);
    chat.append(sendBtn);
  }
  //get senderId, receiverId, msg
  async init() {
    //DRY
    setInterval(() => {
      console.log("call set interva ");
      this.ws.send(JSON.stringify({ sender: "", type: "getusers" }));
    }, 9000);

    let newuser = {
      sender: super.getUserSession(),
      type: "newuser",
    };
    //client 1 enter chat service ->
    this.ws.onopen = () => this.ws.send(JSON.stringify(newuser));

    this.ws.onmessage = (e) => {
      let chatBox = document.getElementById("chatbox").contentDocument;
      let text = "";
      let msg = JSON.parse(e.data);
      // let time = new Date(msg.date);
      // let timeStr = time.toLocaleTimeString();
      switch (msg.type) {
        case "listusers":
          this.showOnlineUsers(msg.users);
          break;
        case "listmessages":
          // document.getElementById("chatbox").innerHTML = ''
          document.getElementById("chatbox").style.display = "block";
          this.showListMessage(msg.messages);
          break;
        case "lastmessage":
          let chatbox = document.getElementById("chatbox");
          let span = document.createElement("span");
          let div = document.createElement("div");
          console.log(msg);
          span.textContent = ` ${msg.message.content} ${msg.message.senttime} ${msg.name} `;
          div.append(span);
          if (chatbox.children != undefined) {
            if (chatbox.children.length > 2) {
              chatbox.children[chatbox.children.length - 3].append(div);
            } else {
              chatbox.children[chatbox.children.length - 2].append(div);
            }
          }
          document.getElementById("messageFieldId").value = "";
          break;
        case "nomessages":
          let chat = document.querySelector("#chatbox");
          chat.innerHTML = "";

          //now no messages -> fix, show message field
          this.showChatWindow(this, msg.receiver);
          super.showNotify("now no messages", "error");
          break;
      }

      if (text.length) {
        chatBox.write(text);
        document.getElementById("chatbox").contentWindow.scrollByPages(1);
      }
      // sjhow sent time, name, fix first create room, added user NOW show another cspanents

      // setInterval(() => {
      //   this.showOnlineUsers(JSON.parse(e.data));
      // }, 5000);

      this.ws.onclose = function (event) {
        if (event.wasClean) {
          console.log("Обрыв соединения"); // например, "убит" процесс сервера
        }
        console.log("Код: " + event.code + " причина: " + event.reason);
      };

      this.ws.onerror = function (error) {
        console.log("Ошибка " + error.message);
      };
    };
  }

  async getHtml() {
    // /?DRY
    //show online user & sended message - like history users
    //listUser - dynamic, create history window, textarea, and btn -> history dynamic change data
    let body = `
    <div id="userlistbox" > </div>
      <div style="display:none" id="chatbox" class="chat_container" >  </div>
      <div id="message_container" >
      </div>
    `;
    return super.showHeader() + body;
  }
}
