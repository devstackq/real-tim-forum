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

  //fix show message - case nomessages UserName
  // fix - updated message from anther client - lastmessage
  // fix-> show : receipment  name

  //fix - send message -listUsers , show Receipment Name fixed, div-inside, div, each time, no delete Receip name, and clear another data
  //fix  todo : lastmessage - listusers - lastmessage show, & add chat with Username
  async getHtml() {
    //<div id="countusers"> </div>
    let body = `
    <div class="chat_wrapper">
    <div id="userlistbox">  </div>
    <div style='display:none' id="message_container">  
    <div id="chatbox" class="chat_container" >      </div>
    <textarea cols="10" rows="10" id="messageFieldId"> </textarea>
    <button id="sendBtnId"> Send message </button
      </div>
      </div>`;
    let header = super.showHeader();
    return header + body;
  }
}
// peredelat pod struct -store
export const setLastMessage = (message, time, senderName, receiver) => {
  let currentChat = document.getElementById(`${receiver}`);
  currentChat.innerHTML = "";
  currentChat.textContent = `Fromz : ${senderName} \n Message: ${message} \n Time:${time} Chat with: `;
};

export const showListMessages = (messages, userid, session, authorName) => {
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
      authorName
    );
  }
};

// update lastmessage & receive message - text format
export const sendMessage = (receiver, authorId, authorName) => {
  let chatContainer = document.querySelector("#message_container");
  let content = chatContainer.children["messageFieldId"].value;
  chatContainer.children["chatbox"].style.display = "block";

  let message = {
    receiver: receiver,
    userid: parseInt(authorId),
    type: "newmessage",
    content: content,
  };

  let div = document.createElement("div");
  let span = document.createElement("span");
  span.className = "chat_sender";

  span.textContent = `${authorName} :  \n${
    message.content
  }   ${new Date().toLocaleTimeString()}  `;

  //current chat -> update contetn
  setLastMessage(
    content,
    new Date().toLocaleTimeString(),
    authorName,
    receiver
  );
  div.append(span);
  chatContainer.children["chatbox"].append(div);
  chatContainer.children["messageFieldId"].value = "";

  toggleOnlineUser(receiver, "prepend");

  wsConn.send(JSON.stringify(message));
};
