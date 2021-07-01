import Parent from "./Parent.js";

export default class Chat extends Parent {
  constructor() {
    super();
  }

  setTitle(title) {
    document.title = title;
  }

  //get senderId, receiverId, msg
  async init() {
    console.log("chat ");
    //DRY

    let ws = new WebSocket("ws://localhost:6969/api/chat");
    console.log(ws);
    ws.addEventListener("message", (e) => {
      console.log(JSON.parse(e.data), "get data from back ws");
    });
    //input name, message current user
    let obj = { receiverId: 0, message: "" };
    obj.receiverId = 2;
    obj.message = "hello dream team !";

    //check state -> then send message
    ws.onopen = () => ws.send(JSON.stringify(obj));

    ws.onclose = function (event) {
      if (event.wasClean) {
        console.log("Соединение закрыто чисто");
      } else {
        console.log("Обрыв соединения"); // например, "убит" процесс сервера
      }
      console.log("Код: " + event.code + " причина: " + event.reason);
    };

    ws.onmessage = function (event) {
      console.log("Получены данные " + event.data);
    };

    ws.onerror = function (error) {
      console.log("Ошибка " + error.message);
    };

    console.log("send object ws ");
    //getLisrUser() & online and offline
    //click -> userId -> getHistoryByChatId()
    //click -> send msg -> webws -> save msg, notify another user
  }

  async getHtml() {
    // /?DRY
    //show online user & sended message - like history users
    //listUser - dynamic, create history window, textarea, and btn -> history dynamic change data
    let body = `
      <div id="listUser" > list users: </div>
      <div id="chat" >message users </div>
      <div id="message_container" >
      <textarea  id="message"> </textarea> 
      <button id="sendMessage" > send </button>
      </div>
    `;
    return super.showHeader() + body;
  }
}
