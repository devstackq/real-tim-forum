import Parent from "./Parent.js";

export default class Chat extends Parent {
  constructor() {
    super();
    this.ws = new WebSocket("ws://localhost:6969/api/chat");
    this.users = {};
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
          let chatWindow = document.getElementById("chatbox");
          chatWindow.innerHTML = "";
          chatWindow.style.display = "block";
          // let userid = super.getUserId();
          let obj = {
            receiver: k,
            sender: super.getUserSession(),
            type: "getmessages",
          };
          this.ws.send(JSON.stringify(obj));
        };

        li.textContent = v;
        // li += v + "<br>"; innerHtml
        ul.append(li);
      }
    }
    document.getElementById("userlistbox").append(ul);
  }

  sendMessage(receiver) {
    let message = {
      content: document.getElementById("messageFieldId").value,
      sender: super.getUserSession(),
      receiver: receiver,
      type: "newmessage",
    };
    this.ws.send(JSON.stringify(message));
  }
  //DRY ?
  showListMessage(messages) {
    let userid = super.getUserId();
    if (messages != null) {
      let chat = document.querySelector("#chatbox");
      let textarea = document.createElement("textarea");
      let sendBtn = document.createElement("button");
      textarea.id = "messageFieldId";
      sendBtn.id = "sendBtnId";
      sendBtn.textContent = "send message";

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

      sendBtn.onclick = this.sendMessage.bind(this, messages[0]["receiver"]);
      chat.append(textarea);
      chat.append(sendBtn);
    } else {
      console.log("no msg");
      super.showNotify("no have message", "error");
    }
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
      let time = new Date(msg.date);
      let timeStr = time.toLocaleTimeString();

      switch (msg.type) {
        case "listusers":
          this.showOnlineUsers(msg.users);
          break;
        case "messagesreceive":
          //why claaed ? now type l;istusers
          console.log(msg);
          this.showListMessage(msg.messages);
          break;
        case "receivemessage":
          // this.showListMessage(msg.messages);
          console.log("receivemessage", msg);

          // append current chatbox, newmessage
          let chatbox = document.getElementById("chatbox");
          //get chatbox children, appen last item
          let li = document.createElement("li");
          li.textContent = msg.content;
          chatbox.children[chatbox.children.length - 3].append(li);
          break;
      }

      if (text.length) {
        chatBox.write(text);
        document.getElementById("chatbox").contentWindow.scrollByPages(1);
      }

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
      <div id="chatbox" class="chat_container" >  </div>
      <div id="message_container" >
      </div>
    `;
    return super.showHeader() + body;
  }
}
