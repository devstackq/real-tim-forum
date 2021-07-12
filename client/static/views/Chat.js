import Parent from "./Parent.js";

export default class Chat extends Parent {
  constructor() {
    super();
    this.ws = super.getWebsocket();
    this.users = [];
    this.onlineUsers = new Map();
    this.historyUsers = [];
    this.chatbox = document.getElementById("chatbox");
    this.HtmlElems = {
      messageContainer: null,
    };
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
    let ul = document.getElementById("listusersID");
    ul.innerHTML = "";
    console.log(Object.entries(object));

    // iter obj users
    for (let [k, user] of Object.entries(object)) {
      this.users.push(user);
      this.onlineUsers.set(key, user);
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
            //remove prev clicked elem class
            for (let i = 0; i < ul.children.length; i++) {
              if (ul.children[i].className == "current") {
                ul.children[i].classList.remove("current");
              }
            }
            li.className = "current";
            // this.HtmlElems.messageContainer.children["chatbox"].style.display =
            //   "none";
            // chatbox.innerHTML = "";
            let obj = {
              receiver: user["UUID"],
              sender: super.getUserSession(),
              type: "getmessages",
            };
            this.ws.send(JSON.stringify(obj));
          };
        }
        // console.log(user['UUID'], super.getUserSession())

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
    let uid = super.getUserId();
    let content =
      this.HtmlElems.messageContainer.children["messageFieldId"].value;
    this.HtmlElems.messageContainer.children["chatbox"].style.display = "block";
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
      if (v["UUID"] === senderUUID) {
        senderName = v.fullname;
      }
    }
    let div = document.createElement("div");
    let span = document.createElement("span");
    span.className = "chat_sender";

    span.textContent = `${senderName} :  \n
     ${message.content}   ${new Date().toLocaleTimeString()}  `;

    div.append(span);
    this.HtmlElems.messageContainer.children["chatbox"].append(div);
    this.HtmlElems.messageContainer.children["messageFieldId"].value = "";
    this.ws.send(JSON.stringify(message));
  }
  //DRY ?
  showListMessage(messages) {
    let userid = super.getUserId();
    if (messages != null) {
      this.HtmlElems.messageContainer.style.display = "block";
      this.HtmlElems.messageContainer.children["chatbox"].style.display =
        "block";
      this.HtmlElems.messageContainer.children["chatbox"].innerHTML = "";

      messages.forEach((item) => {
        let div = document.createElement("div");
        for (let [k, v] of Object.entries(item)) {
          let span = document.createElement("span");

          if (k == "aname" || k == "senttime" || k == "content") {
            span.textContent = `${k == "aname" ? v : ""}  ${
              k == "content" ? v : ""
            }  ${k == "senttime" ? v : ""} \n `;
          }
          if (k == "userid") {
            if (v == userid) {
              div.classList.add("chat_sender");
            }
          }
          div.append(span);
        }
        this.HtmlElems.messageContainer.children["chatbox"].append(div);
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
    //dry?

    this.HtmlElems.messageContainer.children["sendBtnId"].onclick =
      this.sendMessage.bind(this, receiver);
  }

  //get senderId, receiverId, msg
  async init() {
    this.HtmlElems.messageContainer =
      document.querySelector("#message_container");
    //DRY
    // let newuser = {
    //   sender: super.getUserSession(),
    //   type: "newuser",
    // };
    // //client 1 enter chat service ->
    // user signin - profile -> add user -> /chat -> getListuser
    // this.ws.open = () => {
    this.ws.send({ type: "handshake" });
    // };
    console.log(this.ws.readyState, this.ws);
    let n = new WebSocket("ws://localhost:6969/api/chat");
    n.open = () => {
      n.send(JSON.stringify({ type: "getusers" }));
    };

    console.log(n.readyState);

    n.onmessage = (e) => {
      let text = "";
      let msg = JSON.parse(e.data);
      // let time = new Date(msg.date);
      // let timeStr = time.toLocaleTimeString();
      switch (msg.type) {
        case "listusers":
          this.showOnlineUsers(msg.users);
          break;
        case "listmessages":
          document.getElementById("notify").value = "";
          this.HtmlElems.messageContainer.children["chatbox"].style.display =
            "block";
          this.showListMessage(msg.messages);
          break;
        case "nomessages":
          this.HtmlElems.messageContainer.children["chatbox"].style.display =
            "none";
          this.HtmlElems.messageContainer.style.display = "block";
          this.HtmlElems.messageContainer.children["chatbox"].innerHTML = "";
          //now no messages -> fix, show message field
          this.showChatWindow(this, msg.receiver);
          super.showNotify("now no messages", "error");
          break;

        case "lastmessage":
          let span = document.createElement("span");
          let div = document.createElement("div");
          //dry
          // fix receive - sender -> message -> correct show name
          //chat setInterval work only if -> /router -> caht
          //           left join - refactor
          // fix - another user choice 1 ->hide / show
          //how much time call create db ?
          span.textContent = `${msg.message.aname} : \n ${msg.message.content} ${msg.message.senttime}  `;
          div.append(span);
          this.HtmlElems.messageContainer.children["chatbox"].append(div);
          this.HtmlElems.messageContainer.children["messageFieldId"].value = "";
          break;
        case "leave":
          this.onlineUsers.delete(msg.receiver);
          console.log(this.onlineUsers);
          console.log(msg, "leave user");
          break;
      }

      if (text.length) {
        // chatBox.write(text);
        document.getElementById("chatbox").contentWindow.scrollByPages(1);
      }
      // sjhow sent time, name, fix first create room, added user NOW show another cspanents

      this.ws.onclose = function (event) {
        // if (event.wasClean) {
        // console.log("Обрыв соединения"); // например, "убит" процесс сервера
        // }
        console.log("Обрыв соединения"); // например, "убит" процесс сервера
        // this.ws.send(JSON.stringify(obj));
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
    <div id="userlistbox" > <ul id="listusersID" > </ul> </div>
    <div style='display:none' id="message_container"  >  
    <div style='display:none' id="chatbox" class="chat_container" >      </div>
    <textarea id="messageFieldId"> </textarea>
    <button id="sendBtnId"> Send message </button
      </div>`;
    return super.showHeader() + body;
  }
}
