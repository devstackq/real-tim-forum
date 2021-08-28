import Parent from "./Parent.js";

import { wsInit, wsConn, chatStore, getCookie } from "./WebSocket.js";
import { toggleOnlineUser, updateDataInListUser } from "./HandleUsers.js";

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
    <div id="userlistbox"> </div>
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

const debounce = (func, timeout) => {
  let timer;
  return (...args) => {
    clearTimeout(timer);
    timer = setTimeout(() => {
      func.apply(this, args);
    }, timeout);
  };
};

export const appendLastMessageInActiveChat = (type, name, time, content) => {
  let chatContainer = document.querySelector("#message_container");
  chatContainer.children["chatbox"].style.display = "block";
  let activeChat = chatContainer.children["chatbox"];

  let div = document.createElement("div");
  let span = document.createElement("span");
  span.style.padding = "9px";
  if (type == "send") {
    span.className = "chat_sender";
  } else if (type == "receive") {
    span.className = "chat_receiver";
  }
  span.textContent = `${name} ${content} ${time}`;

  div.append(span);
  activeChat.append(div);
  //set last item inside chat window
  activeChat.children[activeChat.children.length - 1].scrollIntoView();
  //empty textarea field
  chatContainer.children["messageFieldId"].value = "";
};

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
        div.classList.add("chat_receiver");
      }
      div.append(span);
      chatContainer.children["chatbox"].append(div);
    });
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
  if (localStorage.getItem("isAuth") == "true") {
    let content =
      document.querySelector("#message_container").children["messageFieldId"]
        .value;
    if (content != "" && content != 0 && content.length > 0) {
      //append last message in activeChat
      appendLastMessageInActiveChat(
        "send",
        authorName,
        new Date().toLocaleTimeString(),
        content
      );
      //update value in list users
      updateDataInListUser(
        receiver,
        new Date().toLocaleTimeString(),
        authorName,
        content
      );
      //change online user position in list users or chane state - online offline
      toggleOnlineUser(receiver, "prepend");
      //send ws  msg
      let message = {
        receiver: receiver,
        userid: parseInt(authorId),
        type: "newmessage",
        content: content,
      };
      wsConn.send(JSON.stringify(message));
    } else {
      alert("empty value");
    }
  } else {
    window.location.replace("/signin");
  }
};
