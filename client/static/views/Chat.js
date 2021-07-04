import Parent from "./Parent.js";

export default class Chat extends Parent {
  constructor() {
    super();
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
  async openChat() {
    let chatWindow = document.getElementById("chat");
    chatWindow.style.display = "block";
    let currentUserUuid = super.getUserSession();
    let userid = super.getUserId();

    let obj = { receiver: this.value, sender: currentUserUuid };
    let response = await super.fetch("chat/history", obj);
    // console.log(response);
    if (response != null) {
      // this.render(response, "#chat");
      response.forEach((item) => {
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
      // chatWindow.textContent = res;
    } else {
      super.showNotify("no have message", "error");
    }
  }

  showOnlineUsers(users) {
    let list = document.getElementById("listUser");
    list.innerHTML = "";
    for (let [k, v] of Object.entries(users)) {
      if (Object.entries(users).length == 1) {
        super.showNotify("Now, no has online users", "error");
        return;
      }
      if (k != super.getUserSession() && Object.entries(users).length > 1) {
        let p = document.createElement("p");

        p.value = k;
        p.onclick = this.openChat;
        p.textContent = v;
        list.append(p);
      }
    }
  }
  //get senderId, receiverId, msg
  async init() {
    //DRY
    // if (wsusers.readyState !== 1) {

    //1 websocket - all, type other
    //if addnewUser -> send client  updated list server

    //best practice?
    // setInterval(() => {
    let wsusers = new WebSocket("ws://localhost:6969/api/chat/users");
    let list = { type: "listusers" };
    wsusers.onopen = () => wsusers.send(JSON.stringify(list));
    wsusers.onmessage = (e) => this.showOnlineUsers(JSON.parse(e.data));

    //onerror    // window.location.replace("/signin");

    //add newuser
    let ws = new WebSocket("ws://localhost:6969/api/chat/newuser");
    ws.addEventListener("message", (e) => {
      console.log(JSON.parse(e.data), "get data from back ws");
    });
    //check state -> then send message
    ws.onopen = () => ws.send(JSON.stringify({ type: "newuser" }));

    ws.onclose = function (event) {
      if (event.wasClean) {
        console.log("Соединение закрыто чисто");
      } else {
        console.log("Обрыв соединения"); // например, "убит" процесс сервера
      }
      console.log("Код: " + event.code + " причина: " + event.reason);
    };
    ws.onmessage = function (event) {
      console.log("Получены данные " + event.data, JSON.parse(event.data));
    };
    ws.onerror = function (error) {
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
