import Parent from "./Parent.js";

export default class Chat extends Parent {
  constructor() {
    super();
    this.ws = new WebSocket("ws://localhost:6969/api/chat");
    this.users = [];
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

  showOnlineUsers(users) {
    //filter if have new user -> append else nothing
    let parent = document.getElementById("userlistbox");
    let ul = document.createElement("ul");
    parent.innerHTML = "";
    ul.id = "listusersID";
    // for (let [k, user] of Object.entries(object)) {
    //1 time push -> ourself
    // this.users.push(users[0]);
    let li = "";
    if (users.length == 1) {
      super.showNotify("Now, no has online user", "error");
      return;
    }
    users.forEach((user) => {
      console.log(user, 2);

      if (user.fullname == "fullname") {
        li = document.createElement("li");
        li.textContent = user.fullname;
        //dry
        //todo: //change color if clicked User current
        let listUsers = document.getElementById("userlistbox").children;
        console.log(listUsers, 1);
        for (let i = 0; i < listUsers.length; i++) {
          if (listUsers[i].className == "current") {
            // console.log(listUsers[(i, 2)], 22);
            listUsers[i].classList.delete("current");
          }
        }

        li.onclick = (e) => {
          // li.style.backgroundColor = "#fff";
          //prev user -> delete class
          li.classList.add("current");
          // this.HtmlElems.messageFieldId.style.display='block'
          this.HtmlElems.messageContainer.style.display = "block";
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

      // if (Object.entries(object).length == 1) {
      //   super.showNotify("no has online user", "error");
      // }
      // // }
    });

    parent.append(ul);
  }

  sendMessage(receiver) {
    let uid = super.getUserId();
    // let content = document.getElementById("messageFieldId").value;
    let content =
      this.HtmlElems.messageContainer.children["messageFieldId"].value;
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
        senderName = v.aname;
      }
    }
    let div = document.createElement("div");
    let span = document.createElement("span");
    span.className = "chat_sender";

    // span.textContent = `${msg.aname} : \n ${msg.message.content} ${msg.message.senttime}  `;

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
    // this.HtmlElems.messageContainer.append(chatbox);
  }

  //get senderId, receiverId, msg
  async init() {
    this.HtmlElems.messageContainer =
      document.querySelector("#message_container");
    // this.HtmlElems.chatBox = document.getElementById("chatbox");
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
      let text = "";
      let msg = JSON.parse(e.data);
      // let time = new Date(msg.date);
      // let timeStr = time.toLocaleTimeString();
      switch (msg.type) {
        case "listusers":
          //filter if have new user -> append else nothing || setInterval pause / play - if length users change
          // for (let i = 0; i < this.users.length; i++) {
          //   console.log(this.users[i]);
          // }
          //todo:
          insert new user -> if no have this.useres
          ds
          // console.log(msg.users, "21", Object.entries(msg.users).length);
          if (Object.entries(msg.users).length > 1) {
            console.log("len ==2", this.users);
            // for (let [key, object] of Object.entries(msg.users)) {
            for (let [wsKey, wsUser] of Object.entries(msg.users)) {
              console.log(wsUser["UUID"], 8, this.users);
              this.users.forEach((user) => {
                console.log(user, 9);
                if (user["UUID"] !== wsUser["UUID"]) {
                  //append new user
                  this.users.push(wsUser);
                  this.showOnlineUsers(this.users);
                }
              });
              // }
            }
          } else {
            this.users.push(msg.users);
            this.showOnlineUsers(this.users);
          }
          break;
        case "listmessages":
          this.HtmlElems.messageContainer.style.display = "block";

          document.getElementById("notify").value = "";
          this.showListMessage(msg.messages);
          break;
        case "lastmessage":
          let span = document.createElement("span");
          let div = document.createElement("div");
          //dry
          this.HtmlElems.messageContainer.children["chatbox"].style.display =
            "block";
          for (let [k, v] of Object.entries(this.users)) {
            if (v["UUID"] === msg.receiver) {
              msg.aname = v.aname;
            }
          }
          span.textContent = `${msg.aname} : \n ${msg.message.content} ${msg.message.senttime}  `;
          div.append(span);

          this.HtmlElems.messageContainer.children["chatbox"].append(div);
          this.HtmlElems.messageContainer.children["messageFieldId"].value = "";
          break;
        case "nomessages":
          this.HtmlElems.messageContainer.children["chatbox"].style.display =
            "none";
          this.HtmlElems.messageContainer.children["chatbox"].innerHTML = "";
          //now no messages -> fix, show message field
          this.showChatWindow(this, msg.receiver);
          super.showNotify("now no messages", "error");
          break;
      }

      if (text.length) {
        // chatBox.write(text);
        document.getElementById("chatbox").contentWindow.scrollByPages(1);
      }
      // sjhow sent time, name, fix first create room, added user NOW show another cspanents
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
    <div id="message_container" style='display:none' >  
    <div id="chatbox" class="chat_container" >      </div>
    <textarea id="messageFieldId"> </textarea>
    <button id="sendBtnId"> Send message </button
      </div>`;
    return super.showHeader() + body;
  }
}
