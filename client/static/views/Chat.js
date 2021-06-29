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
    let response = await super.fetch("chat", null);
    //getLisrUser() & online and offline
    //click -> userId -> getHistoryByChatId()
    //click -> send msg -> websocket -> save msg, notify another user
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
      
    `
    return super.showHeader() + body
  }
}
