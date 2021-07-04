import WebSocket from "./WebScoket.js";

export default class Chat extends WebSocket {
  // setWS(value) {
  //   this.ws.conn = value;
  // }
  constructor(ws) {
    super();
    this.ws = {
      conn: new WebSocket("ws://localhost:6969/api/chat"),
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

  openChat() {
    let chatWindow = document.getElementById("chat");
    chatWindow.innerHTML = "";
    chatWindow.style.display = "block";
    let currentUserUuid = super.getUserSession();
    let userid = super.getUserId();

    let obj = {
      receiver: this.getAttribute("uuid"),
      sender: currentUserUuid,
      type: "getmessages",
    };
    let ws = super.getWs();
    // let ws = new WebSocket("ws://localhost:6969/api/chat");
    // let ws = this.value;
    ws.onopen = () => ws.send(JSON.stringify(obj));

    ws.onmessage = (e) => {
      // this.render(response, "#chat");
      JSON.parse(e.data).forEach((item) => {
        let div = document.createElement("div");
        for (let [k, v] of Object.entries(item)) {
          let span = document.createElement("span");
          span.textContent = ` ${k} : ${v} `;
          if (k == "userid") {
            if (v == userid) {
              div.classList.add("chat_sender");
            }
          }
          div.append(span);
        }
        chatWindow.append(div);
      });
      // } else {
      //   super.showNotify("no have message", "error");
    };

    ws.onerror = function (error) {
      console.log("Ошибка " + error.message);
    };

    ws.onclose = function (event) {
      if (event.wasClean) {
        console.log("Соединение закрыто чисто");
      } else {
        console.log("Обрыв соединения"); // например, "убит" процесс сервера
      }
      console.log("Код: " + event.code + " причина: " + event.reason);
    };
  }

  showOnlineUsers(users) {
    // console.log(users);
    let list = document.getElementById("listUser");
    list.innerHTML = "";
    for (let [k, v] of Object.entries(users)) {
      if (Object.entries(users).length == 1) {
        super.showNotify("Now, no has online users", "error");
        return;
      }
      if (k != super.getUserSession() && Object.entries(users).length > 1) {
        let p = document.createElement("p");
        p.value = this.getWS();
        p.setAttribute("uuid", k);
        p.onclick = this.openChat;
        p.textContent = v;
        list.append(p);
      }
    }
  }
  sendMessage() {
    console.log("sed");
  }
  //get senderId, receiverId, msg
  async init() {
    //DRY
    document.getElementById("message_container").onclick = this.sendMessage;
    // this.value = this.ws.conn;
    // if (wsusers.readyState !== 1) {
    //1 websocket - all, type other
    //if addnewUser -> send client  updated list server

    //best practice?
    // setInterval(() => {
    //onerror    // window.location.replace("/sigWebSocket { url: "ws://localhost:6969/
    //newuser
    // this.setWS(new WebSocket("ws://localhost:6969/api/chat"));
    this.ws.conn.onopen = () =>
      this.ws.conn.send(JSON.stringify({ type: "newuser" }));
    //list online users
    // this.ws.conn.onopen = () => this.ws.conn.send(JSON.stringify({ type: "listusers" }));

    // if(JSON.parse(e.data == "newuser")
    this.ws.conn.onmessage = (e) => {
      setInterval(() => {
        this.showOnlineUsers(JSON.parse(e.data));
      }, 5000);
    };

    this.ws.conn.onclose = function (event) {
      if (event.wasClean) {
        console.log("Обрыв соединения"); // например, "убит" процесс сервера
      }
      console.log("Код: " + event.code + " причина: " + event.reason);
    };

    this.ws.conn.onerror = function (error) {
      console.log("Ошибка " + error.message);
    };

    //data list user -> send  by uuid
    // this.showOnlineUsers(JSON.parse(event.data));

    //getLisrUser() & online and offline
    //click -> userId -> getHistoryByChatId()
    //click -> send msg -> webws -> save msg, notify another user
  }

  async getHtml() {
    // /?DRY
    //show online user & sended message - like history users
    //listUser - dynamic, create history window, textarea, and btn -> history dynamic change data
    let body = `
    <div id="listUser" > </div>
      <div id="chat" class="chat_container" >  </div>
      <div id="message_container" >
      <textarea  id="message"> </textarea> 
      <button id="sendMessage" > send </button>
      </div>
    `;
    return super.showHeader() + body;
  }
}
