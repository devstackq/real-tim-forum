import Parent from "./Parent.js";
import { wsInit, wsConn, toggleOnlineUser } from "./WebSocket.js";

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
    wsInit("", "getusers"); //open conn ?
  }

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
// export const setLastMessage = (
//   message,
//   time,
//   senderName,
//   receiver,
//   receiverName
// ) => {
//   let currentChat = document.getElementById(`${receiver}`);
//   currentChat.innerHTML = "";
//   currentChat.textContent = `Fromz : ${senderName} \n Message: ${message} \n Time:${time} Chat with:  ${receiverName}`;
// };

export const showListMessages = (messages, userid, session, authorName) => {
  let chatContainer = document.querySelector("#message_container");
  if (messages != null && chatContainer != null) {
    chatContainer.style.display = "block";
    chatContainer.children["chatbox"].style.display = "block";
    chatContainer.children["chatbox"].innerHTML = "";
    // console.log(messages, 321321);
    let receivername = "";

    messages.forEach((item) => {
      // console.log(item);
      let div = document.createElement("div");
      let span = document.createElement("span");
      span.textContent = `${item.sendername}  ${item.content} ${item.senttime} \n `;
      if (item.userid == userid) {
        div.classList.add("chat_sender");
      }
      div.append(span);
      chatContainer.children["chatbox"].append(div);
      if (item.sendername != authorName) {
        receivername = item.sendername;
      }
    });
    //call func
    let receive = "";
    if (messages.length != 0) {
      receive = messages[0]["receiver"];
    }
    // console.log(receive, "list mes reciece");
    toggleOnlineUser(receive);

    chatContainer.children["sendBtnId"].onclick = sendMessage.bind(
      this,
      receive,
      userid,
      authorName,
      session
    );
  }
};

// update lastmessage & receive message - text format
export const sendMessage = (receiver, authorId, authorName, session) => {
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
  //dry
  let el = document.getElementById(receiver);
  el.textContent = ` ${authorName} ${content} ${new Date().toLocaleTimeString()}`;

  div.append(span);
  chatContainer.children["chatbox"].append(div);
  chatContainer.children["messageFieldId"].value = "";

  toggleOnlineUser(receiver, "prepend");

  wsConn.send(JSON.stringify(message));
};
//test listuser, send msg, signin, signup, show msg, receive msg
