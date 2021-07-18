import { showListUser, addNewUser, listUsers } from "./HandleUsers.js";
import { showListMessages, sendMessage } from "./Chat.js";

export let wsConn = null;

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

// add user - send uuid
export const wsInit = (...args) => {
  if (wsConn == null) {
    console.log(wsConn, "val, singleton?");
    wsConn = new WebSocket("ws://localhost:6969/api/chat");
    addNewUser(args[0]);
  }
  let chatContainer = document.querySelector("#message_container");

  wsConn.onmessage = (e) => {
    let message = JSON.parse(e.data);
    console.log(message.type, typeof message.users, message.users);
    // listUsers = message.users;
    switch (message.type) {
      case "observeusers":
        console.log(message.timeusers, message.alphausers);
        //show divs class users, inside each user - - div
        //  show by reverse
        //update all conn -> added new user
        // for (let i = 0; i < message.users.length; i++) {}
        showListUser(message.users);
        break;
      case "getusers":
        //add in arrayUsers
        for (let i = 0; i < message.users.length; i++) {}
        if (listUsers.length == 0) {
          listUsers.push(message.users);
        }
        //work this var
        console.log(listUsers, "listtt");
        showListUser(message.users, "getusers");
        break;
      case "listmessages":
        document.getElementById("notify").value = "";
        // chatContainer.children["chatbox"].style.display =
        //   "block";
        showListMessages(
          message.messages,
          getUserId(),
          getCookie("session"),
          message.users
        );
        break;
      case "nomessages":
        alert("no messages now..");
        document.getElementById("notify").value = "no message now...";
        // chatContainer.children["chatbox"].style.display = "none";
        chatContainer.style.display = "block";
        chatContainer.children["chatbox"].innerHTML = "";
        //now no messages -> fix, show message field
        chatContainer.children["sendBtnId"].onclick = sendMessage.bind(
          this,
          message.receiver,
          getUserId(),
          getCookie("session"),
          message.users
        );
        // super.showNotify("now no messages", "error");
        break;
      case "lastmessage":
        let span = document.createElement("span");
        let div = document.createElement("div");
        chatContainer.children["chatbox"].style.display = "block";
        span.textContent = `${message.message.aname} : \n ${message.message.content} ${message.message.senttime}  `;
        div.append(span);
        chatContainer.children["chatbox"].append(div);
        chatContainer.children["messageFieldId"].value = "";
        break;
      case "leave":
        // this.onlineUsers.delete(message.receiver);
        showListUser(message.users);
        break;
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
