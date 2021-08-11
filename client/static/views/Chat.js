import Parent from "./Parent.js";

import { wsInit, wsConn, chatStore, getCookie } from "./WebSocket.js";
import { toggleOnlineUser } from "./HandleUsers.js";

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
  // chatDiv.removeEventListener("scroll", debounce, false);

  async init() {
    this.HtmlElems.messageContainer =
      document.querySelector("#message_container");
    wsInit(getCookie("session"), "getusers"); //open conn ?

    let chatDiv =
      document.querySelector("#message_container").children["chatbox"];
    //scroll, each 240 ms, call anonim func, if offsetTop ==0 -> get ws data
    chatDiv.addEventListener(
      "scroll",
      debounce(() => {
        if (chatStore.messageLen > 9) {
          // console.log(countNewMessage.value, "count");
          if (chatDiv.scrollTop <= 1) {
            let obj = {
              receiver: chatStore.receiver,
              sender: chatStore.sender,
              type: "last10msg",
              offset: chatStore.offset + 10 + chatStore.countNewMessage,
            };
            wsConn.send(JSON.stringify(obj));
          }
        }
      }, 240)
    );
  }

  async getHtml() {
    //<div id="countusers"> </div>
    let body = `
    <div class="chat_wrapper">
    <div id="userlistbox">  </div>
    <div style='display:none' id="message_container">  
    <div id="chatbox" class="chat_container">      </div>
    <textarea cols="10" rows="10" id="messageFieldId"> </textarea>
    <button id="sendBtnId"> Send message </button
      </div>
      </div>`;
    let header = super.showHeader();
    return header + body;
  }
}

function debounce(func, timeout) {
  let timer;
  return (...args) => {
    clearTimeout(timer);
    timer = setTimeout(() => {
      func.apply(this, args);
    }, timeout);
  };
}

export const showListMessages = (messages, userid, session, authorName) => {
  let chatContainer = document.querySelector("#message_container");
  if (messages != null && chatContainer != null) {
    chatContainer.style.display = "block";
    chatContainer.children["chatbox"].style.display = "block";
    chatContainer.children["chatbox"].innerHTML = "";

    messages.forEach((item, index) => {
      let div = document.createElement("div");
      let span = document.createElement("span");
      span.textContent = `${item.sendername}  ${item.content} ${item.senttime} \n `;
      if (item.userid == userid) {
        div.classList.add("chat_sender");
      } else {
        // !item.isread  send ws, & remove class unread
        div.classList.add("chat_receiver");
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
  span.style.padding = "9px";
  div.append(span);
  let chatDiv = chatContainer.children["chatbox"];
  chatDiv.append(div);
  //set last item inside chat window
  chatDiv.children[chatDiv.children.length - 1].scrollIntoView();
  chatContainer.children["messageFieldId"].value = "";
  toggleOnlineUser(receiver, "prepend");
  wsConn.send(JSON.stringify(message));
  chatStore.countNewMessage += 1;
};
//test listuser, send msg, signin, signup, show msg, receive msg
