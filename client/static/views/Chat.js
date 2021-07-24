import Parent from "./Parent.js";
import { wsInit, wsConn, getCookie, toggleOnlineUser } from "./WebSocket.js";

export default class Chat extends Parent {
  constructor() {
    super();
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
    //getusers
    if (wsConn != null && wsConn.readyState == 1) {
      wsConn.send(
        JSON.stringify({ sender: getCookie("session"), type: "newuser" })
      );
    }
    // showListUser(ListUsers);
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
// peredelat pod struct -store

export const showListMessages = (messages, userid, session, author) => {
  let chatContainer = document.querySelector("#message_container");
  if (messages != null && chatContainer != null) {
    chatContainer.style.display = "block";
    chatContainer.children["chatbox"].style.display = "block";
    chatContainer.children["chatbox"].innerHTML = "";

    messages.forEach((item) => {
      let div = document.createElement("div");
      let span = document.createElement("span");

      span.textContent = `${item.aname}  ${item.content}  ${item.senttime} \n `;

      if (item.userid == userid) {
        div.classList.add("chat_sender");
      }
      div.append(span);
      chatContainer.children["chatbox"].append(div);
    });
    //call func
    let receive = "";
    if (messages.length != 0) {
      receive = messages[0]["receiver"];
    }

    toggleOnlineUser(receive);

    chatContainer.children["sendBtnId"].onclick = sendMessage.bind(
      this,
      receive,
      userid,
      session,
      author
    );
  }
};

export const sendMessage = (receiver, authorId, senderUUID, author) => {
  let chatContainer = document.querySelector("#message_container");
  let content = chatContainer.children["messageFieldId"].value;
  chatContainer.children["chatbox"].style.display = "block";

  let message = {
    content: content,
    sender: senderUUID,
    receiver: receiver,
    userid: parseInt(authorId),
    type: "newmessage",
  };

  let div = document.createElement("div");
  let span = document.createElement("span");
  span.className = "chat_sender";

  span.textContent = `${author} :  \n${
    message.content
  }   ${new Date().toLocaleTimeString()}  `;

  div.append(span);
  chatContainer.children["chatbox"].append(div);
  chatContainer.children["messageFieldId"].value = "";

  toggleOnlineUser(receiver, "prepend");

  wsConn.send(JSON.stringify(message));
};
