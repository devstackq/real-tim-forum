import Parent from "./Parent.js";

export default class Chat extends Parent {
  constructor() {
    super();
    this.ws = new WebSocket("ws://localhost:6969/api/chat");
    this.users = {};
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

  showOnlineUsers(users) {
    let ul = document.createElement("ul");
    ul.id = "listusersID";
    // console.log(this.ws, 987);
    for (let [k, v] of Object.entries(users)) {
      let li = document.createElement("li");
      if (Object.entries(users).length == 1) {
        super.showNotify("Now, no has online users", "error");
        return;
      }
      if (k != super.getUserSession() && Object.entries(users).length > 1) {
        li.onclick = (e) => {
          // chatbox.innerHTML = "";
          let obj = {
            receiver: k,
            sender: super.getUserSession(),
            type: "getmessages",
          };
          this.ws.send(JSON.stringify(obj));
        };
        li.textContent = v;
        ul.append(li);
      }
    }
    document.getElementById("userlistbox").append(ul);
  }

  sendMessage(receiver) {
    let chatbox = document.querySelector("#chatbox");
    let uid = super.getUserId();
    let content = document.getElementById("messageFieldId").value;

    let message = {
      content: content,
      sender: super.getUserSession(),
      receiver: receiver,
      userid: parseInt(uid),
      type: "newmessage",
    };
    this.ws.send(JSON.stringify(message));
    // append in last item client
    let li = document.createElement("li");
    li.textContent = content;
    li.className = "chat_sender";

    chatbox.children[chatbox.children.length - 2].append(li);
  }
  //DRY ?
  showListMessage(messages) {
    let userid = super.getUserId();
    if (messages != null) {
      let chat = document.querySelector("#chatbox");

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
  //dry
  showChatWindow(scope, receiver) {
    console.log("chat wind call");
    document.getElementById("chatbox").style.display = "block";
    //dry?
    let chat = document.querySelector("#chatbox");
    let textarea = document.createElement("textarea");
    let sendBtn = document.createElement("button");

    textarea.id = "messageFieldId";
    sendBtn.id = "sendBtnId";
    sendBtn.textContent = "send messageq";

    sendBtn.onclick = this.sendMessage.bind(this, receiver);
    chat.append(textarea);
    chat.append(sendBtn);
  }
  //get senderId, receiverId, msg
  async init() {
    //DRY
    let newuser = {
      sender: super.getUserSession(),
      type: "newuser",
    };

    this.ws.onopen = () => this.ws.send(JSON.stringify(newuser));

    this.ws.onmessage = (e) => {
      let chatBox = document.getElementById("chatbox").contentDocument;
      let text = "";
      let msg = JSON.parse(e.data);
      // let time = new Date(msg.date);
      // let timeStr = time.toLocaleTimeString();

      switch (msg.type) {
        case "listusers":
          console.log("add new user");
          //update each 10 sec ? for show new user
          this.showOnlineUsers(msg.users);
          break;
        case "listmessages":
          //get from db all messages
          document.getElementById("chatbox").style.display = "block";
          this.showListMessage(msg.messages);
          break;
        case "lastmessage":
          //dry?
          let chatbox = document.getElementById("chatbox");
          let li = document.createElement("li");
          li.textContent = msg.content;
          chatbox.children[chatbox.children.length - 2].append(li);
          break;
        case "nomessages":
          //now no messages -> fix, show message field
          this.showChatWindow(this, msg.receiver);
          super.showNotify("now no messages", "error");
          break;
      }
      if (text.length) {
        chatBox.write(text);
        document.getElementById("chatbox").contentWindow.scrollByPages(1);
      }
      sjhow sent time, name, fix first create room, added user NOW show another clients

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
