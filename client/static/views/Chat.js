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
    //DRY
    //get user name, getUUID in cookie -> send /ws/chat - then split & use data

    // let response = await fetch("http://localhost:6969/api/profile");
    // console.log(response, "porifle");
    // if (response.status === 200) {
    //   let result = await response.json();
    //   console.log(result.User['fullname'])
    // }
    let obj = { receiverId: "", message: "", type: "newuser" };

    
    let ws = new WebSocket("ws://localhost:6969/api/chat");
    // console.log(ws);
    ws.addEventListener("message", (e) => {
      console.log(JSON.parse(e.data), "get data from back ws");
    });
    //input name, message current user
    // console.log(document.cookie.split("session=")[1].split(";")[0], "uuid")
    // obj.receiver= document.cookie.split("session=")[1].split(";")[0]
    // obj.receiver= click.value()
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
      console.log("Получены данные " + event.data, JSON.parse(event.data));
    };

    ws.onerror = function (error) {
      console.log("Ошибка " + error.message);
    };
//use setTimeout ?
    let wsusers = new WebSocket("ws://localhost:6969/api/getusers");
    let list = {type : "listusers"}
    wsusers.onopen = () => wsusers.send(JSON.stringify(list));

    wsusers.onmessage = function (event) {
      //data list user -> send  by uuid
      console.log("Получены данные list users " + event.data);
    };
// console.log(obj, list)



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
