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
    return super.showHeader();
  }
}
