import { showListUser, addNewUser } from "./HandleUsers.js";
import { showListMessages, sendMessage, setLastMessage } from "./Chat.js";

export let ListUsers = {};

export let wsConn = null;
export let authorName = "";

export function getCookie(cName) {
  const name = cName + "=";
  const cDecoded = decodeURIComponent(document.cookie); //to be careful
  const cArr = cDecoded.split("; ");
  let res;
  cArr.forEach((val) => {
    if (val.indexOf(name) === 0) res = val.substring(name.length);
  });
  return res;
}

export const getUserId = () => {
  if (document.cookie.split(";").length == 3) {
    return document.cookie.split(";")[1].slice(9).toString();
  }
};

export const toggleOnlineUser = (receiver, type) => {
  let currentUser = document.getElementById(receiver);
  let listUsers = document.getElementById("userlistbox"); // out global var ?

  for (let i = 0; i < listUsers.children.length; i++) {
    if (listUsers.children[i].classList.contains("current")) {
      listUsers.children[i].classList.remove("current");
    }
  }
  currentUser.classList.add("current");
  type == "prepend" ? listUsers.prepend(currentUser) : null;
};
// add user - send uuid
export const wsInit = (...args) => {
  if (wsConn == null) {
    console.log(wsConn, "val, singleton?");
    wsConn = new WebSocket("ws://localhost:6969/api/chat");
    addNewUser(args[0]);
  }
  let chatContainer = document.querySelector("#message_container");

  wsConn.onmessage = (e) => {
    let authorId = getCookie("user_id");
    let authorSession = getCookie("session");

    let message = JSON.parse(e.data);
    // listUsers = message.users;
    let el = null;
    if (
      (message.type == "leave" || message.type == "online") &&
      message != null &&
      message.user != undefined
    ) {
      //remove online class - by id or uuid
      el = document.getElementById(message.user.id);
      if (el == null) {
        el = document.getElementById(message.user.uuid);
      }
    }

    switch (message.type) {
      case "online":
        // el == null ? (el = document.getElementById(message.id)) : null;
        el == null ? (el = document.getElementById(message.uuid)) : null;
        el == null ? (el = document.getElementById(message.id)) : null;
        el != null ? (el.className = "online") : null;
        break;
      case "observeusers":
        authorName = message.author;
        showListUser(message.users);
        break;
      case "listmessages":
        document.getElementById("notify").value = "";
        showListMessages(
          message.messages,
          authorId,
          authorSession,
          message.author
        );
        break;
      case "nomessages":
        alert("no messages now..");
        document.getElementById("notify").value = "no message now...";
        // chatContainer.children["chatbox"].style.display = "none";
        chatContainer.style.display = "block";
        chatContainer.children["chatbox"].innerHTML = "";
        //now no messages -> fix, show message field
        console.log(authorSession, message.author, message.receiver, message);
        // authorSession,
        chatContainer.children["sendBtnId"].onclick = sendMessage.bind(
          this,
          message.receiver,
          authorId,
          message.author
        );
        // super.showNotify("now no messages", "error");
        break;
      case "lastmessage":
        //prepend user - first, another client
        chatContainer.children["chatbox"].style.display = "block";
        let span = document.createElement("span");
        span.textContent = `From :${message.message.aname} : \n ${message.message.content} ${message.message.senttime}  Chat with: `;
        chatContainer.children["chatbox"].append(span);
        chatContainer.children["messageFieldId"].value = "";
        // set session || userid
        toggleOnlineUser(message.message.sender, "prepend");
        break;
      case "leave":
        el == null ? (el = document.getElementById(message.uuid)) : null;
        el == null ? (el = document.getElementById(message.id)) : null;
        el != null ? el.classList.remove("online") : null;
        el.id = message.id; //set id, replace - prev uuid
        break;
      default:
        console.log("incorrect type");
    }

    wsConn.onclose = function (event) {
      console.log(
        " Обрыв соединения, Код: " + event.code + " причина: " + event.reason
      );
      wsConn.close();
      // wsConn.send(JSON.stringify({ type: "close" }));
    };
    wsConn.onerror = function (error) {
      console.log("Ошибка " + error.message);
      wsConn.close();
    };
  };
};
