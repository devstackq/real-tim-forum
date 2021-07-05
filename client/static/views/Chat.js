import Parent from "./Parent.js";

export default class Chat extends Parent {
  constructor() {
    super();
    this.ws = new WebSocket("ws://localhost:6969/api/chat");
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

  //DRY
  openChat() {
    console.log("click chat", this);
    let chatWindow = document.getElementById("chatbox");
    chatWindow.innerHTML = "";
    chatWindow.style.display = "block";
    let userid = super.getUserId();

    let obj = {
      receiver: this.getAttribute("uuid"),
      sender: super.getUserSession(),
      type: "listmessages",
    };
    // let ws = this.getAttribute("ws");
    // let ws = this.value;
    // best practice?
    let ws = new WebSocket("ws://localhost:6969/api/chat");

    ws.onopen = () => ws.send(JSON.stringify(obj));

    ws.onmessage = (e) => {
      console.log(JSON.parse(e.data));
      // this.render(response, "#chat");
      JSON.parse(e.data).forEach((item) => {
        let chat = document.querySelector("#chatbox");
        let div = document.createElement("div");

        for (let [k, v] of Object.entries(item)) {
          let span = document.createElement("span");
          if (v != null && v != "") {
            span.textContent = `  ${v}: `;
          }
          if (k == "userid") {
            if (v == userid) {
              div.classList.add("chat_sender");
            }
          }
          // chatBox.write(text);
          div.append(span);
        }
        chat.append(div);
        // document.getElementById("chatbox").contentWindow.scrollByPages(1);
      });
      // } else {
      //   super.showNotify("no have message", "error");
    };
  }

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
        li.setAttribute("uuid", k);
        li.setAttribute("ws", this.ws);
        li.onclick = this.openChat;
        // li.value = this.ws;
        li.textContent = v;
        // li += v + "<br>"; innerHtml
        ul.append(li);
      }
    }
    document.getElementById("userlistbox").append(ul);
  }

  sendMessage() {
    console.log("send");
  }
  //get senderId, receiverId, msg
  async init() {
    //DRY
    // document.onload =
    let newuser = {
      sender: super.getUserSession(),
      type: "newuser",
    };
    this.ws.onopen = () => this.ws.send(JSON.stringify(newuser));

    document.getElementById("message_container").onclick = this.sendMessage;

    // if (wsusers.readyState !== 1) {
    //if addnewUser -> send client  updated list server

    this.ws.onmessage = (e) => {
      let chatBox = document.getElementById("chatbox").contentDocument;
      let text = "";
      let msg = JSON.parse(e.data);
      let time = new Date(msg.date);
      let timeStr = time.toLocaleTimeString();

      switch (msg.type) {
        case "message":
          text =
            "(" + timeStr + ") <b>" + msg.name + "</b>: " + msg.text + "<br>";
          break;
        case "listusers":
          this.showOnlineUsers(msg.users);
        // let ul = document.createElement("ul");
        // for (let [k, v] of Object.entries(msg.users)) {
        //   let li = document.createElement("li");
        //   if (
        //     k != super.getUserSession() &&
        //     Object.entries(msg.users).length > 1
        //   ) {
        //     li.onclick = this.openChat;
        //     li.value = this.ws;
        //     li.textContent = v;
        //     // li += v + "<br>"; innerHtml
        //     ul.append(li);
        //   }
        // }
        // document.getElementById("userlistbox").append(ul);
        // break;
        // case (e) => document.getElementById("listusersID").children.onclick :
        //   console.log(e);
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

      //data list user -> send  by uuid
      // this.showOnlineUsers(JSON.parse(event.data));

      //getLisrUser() & online and offline
      //click -> userId -> getHistoryByChatId()
      //click -> send msg -> webws -> save msg, notify another user
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
      <textarea  id="message"> </textarea> 
      <button id="sendMessage" > send </button>
      </div>
    `;
    return super.showHeader() + body;
  }
}
