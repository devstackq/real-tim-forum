import Parent from "./Parent.js";
import { wsInit, wsConn, getCookie } from "./WebSocket.js";

export default class Chat extends Parent {
  constructor() {
    super();
    this.ws = wsConn;
    this.msgType = "";
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
  async init() {
    this.HtmlElems.messageContainer =
      document.querySelector("#message_container");

    wsInit(); //each time add user
    if (wsConn != null && wsConn.readyState == 1) {
      wsConn.send(
        JSON.stringify({ sender: getCookie("session"), type: "getusers" })
      );
    }
  }

  async getHtml() {
    //?DRY
    let body = `
    <div id="userlistbox" > <ul id="listusersID" > </ul> </div>
    <div style='display:none' id="message_container"  >  
    <div id="chatbox" class="chat_container" >      </div>
    <textarea id="messageFieldId"> </textarea>
    <button id="sendBtnId"> Send message </button
      </div>`;
    return super.showHeader() + body;
  }
}

export const showListMessages = (messages, userid, session, users) => {
  let chatContainer = document.querySelector("#message_container");
  if (messages != null && chatContainer != null) {
    chatContainer.style.display = "block";
    chatContainer.children["chatbox"].style.display = "block";
    chatContainer.children["chatbox"].innerHTML = "";

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
      chatContainer.children["chatbox"].append(div);
    });
    //call func
    let receive = "";
    if (messages.length != 0) {
      receive = messages[0]["receiver"];
    }
    // this.showChatWindow(this, receive);
    chatContainer.children["sendBtnId"].onclick = sendMessage.bind(
      this,
      receive,
      userid,
      session,
      users
    );
  }
};

export const sendMessage = (receiver, userid, senderUUID, users) => {
  let chatContainer = document.querySelector("#message_container");
  let content = chatContainer.children["messageFieldId"].value;
  chatContainer.children["chatbox"].style.display = "block";
  let message = {
    content: content,
    sender: senderUUID,
    receiver: receiver,
    userid: parseInt(userid),
    type: "newmessage",
  };
  let senderName = "";
  for (let [k, v] of Object.entries(users)) {
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
  chatContainer.children["chatbox"].append(div);
  chatContainer.children["messageFieldId"].value = "";
  wsConn.send(JSON.stringify(message));
};
