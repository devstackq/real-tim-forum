import {
  showListUser,
  listUsers,
  toggleOnlineUser,
  updateDataInListUser,
} from "./HandleUsers.js";
import {
  showListMessages,
  sendMessage,
  appendLastMessageInActiveChat,
} from "./Chat.js";

export let chatContainer;
export let chatDiv;
let tempListUsers = [];

export let ListUsers = {};
export let wsConn = null;
export let listMessages = [];
export let countMsgDiv = document.getElementById("unread");

export let chatStore = {
  authorName: "",
  sender: "",
  receiver: "",
  offset: 0,
  messageLen: 0,
  countNewMessage: 0,
};

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
//only update value, switch type:

const appendNewUserInListUsers = (idx, wsMessage, listUsers) => {
  let listUsersDom = document.querySelector("#userlistbox");
  let el = document.createElement("li");
  el.id = wsMessage.user.uuid;
  el.classList.add("online");
  let pattern = `<h3 class="partner">${wsMessage.user.fullname}</h3>
          <span>Now no have messages..</span>
          <span class="time"></span>`;
  el.innerHTML = pattern;

  el.onclick = (e) => {
    //remove prev clicked elem class
    let obj = {
      receiver: wsMessage.user.uuid,
      sender: getCookie("session"),
      type: "last10msg",
    };
    wsConn.send(JSON.stringify(obj));
    chatDiv.textContent = "";
    console.log(wsMessage.user, "append new user func");
    chatDiv.value = wsMessage.user.uuid;
    toggleOnlineUser(wsMessage.user.uuid);
    //set chat - for new user
  };
  listUsersDom.insertBefore(el, listUsersDom.children[idx]);
  //append in array obj - new user
  let temp = [];
  for (let i = 0; i < listUsers.length; i++) {
    temp.push(listUsers[i]);
    if (i == idx) {
      temp.push(wsMessage.user);
    }
  }
  console.log(wsMessage.user, listUsers, idx, "append new usec", temp);
};

//insert new signup user, and sort local array list users
const insertNewUser = (message, tempListUsers) => {
  console.log(message, tempListUsers, "gav");
  if (tempListUsers.length > 1) {
    for (let [index, user] of Object.entries(tempListUsers)) {
      //sort by messages
      if (`${user.lastmessage["String"]}` == "") {
        if (message.user.fullname < user.fullname) {
          appendNewUserInListUsers(index, message, tempListUsers);
          break;
        }
      }
    }
  } else {
    appendNewUserInListUsers(0, message, tempListUsers);
  }
};

// add user - send uuid
export const wsInit = (...args) => {
  if (wsConn == null) {
    wsConn = new WebSocket("ws://localhost:6969/api/chat");
    console.log(wsConn, "val, singleton?");
  }
  if (args[0] != "" && args[1] == "signin") {
    listUsers(args[0], "signin");
  } else if (args[1] == "getusers") {
    listUsers(getCookie("session"), args[1]);
  } else if (args[1] == "newuser" && args[0] != "") {
    listUsers(args[0], args[1]);
  }

  if (args[1] != "newuser" && args[1] != "signin") {
    chatContainer == undefined || chatContainer == null
      ? (chatContainer = document.getElementById("message_container"))
      : null;
    chatDiv == undefined || chatDiv == null
      ? (chatDiv = chatContainer.children["chatbox"])
      : null;
  }

  wsConn.onmessage = (e) => {
    let authorId = getCookie("user_id");
    let authorSession = getCookie("session");
    let message = JSON.parse(e.data);
    let el = null;

    console.log(message.type);

    switch (message.type) {
      case "newuser":
        //signup user, another online user -show him
        //DRY!
        insertNewUser(message, tempListUsers);
        // showListUser(tempListUsers, message, " newuser");
        break;
      case "online":
        //find & replace by id -> uuid
        el == null ? (el = document.getElementById(message.user.id)) : null;
        el == null ? (el = document.getElementById(message.user.uuid)) : null;
        el != null ? el.classList.add("online") : null;
        console.log(message, "ponloien new user");
        el.id = message.user.uuid;
        break;
      case "chatusers":
        //temp, for sort & insert  in DOm, new signup user
        tempListUsers = [];
        tempListUsers = [...message.users];
        showListUser(message.users);
        break;
      case "chathistory":
        // fix - online, offline activeChat
        console.log(message.receiver, chatDiv.value, 111);
        if (message.receiver == chatDiv.value) {
          //prepend reversed get message from backend, offset limit

          listMessages = [...message.messages.reverse(), ...listMessages]; // for compare, & ignoring duplicate msg
          // scroll -> up to 10 MSGesture, position -> send rRequest
          // chatDiv.value = message.receiver;
          //export value use Chat component func
          chatDiv.textContent = "";
          showListMessages(
            listMessages,
            authorId,
            authorSession,
            message.author
          );
          //set userid/uuid - current chat
          chatStore.sender = message.sender;
          chatStore.receiver = message.receiver;
          chatStore.offset = message.offset;
          chatStore.messageLen = message.messages.length;

          chatDiv.children.length > 1
            ? chatDiv.children[chatDiv.children.length - 1].scrollIntoView()
            : null;
        }
        break;
      case "nomessages":
        alert("no have messages..");
        // document.getElementById("notify").value = "no have messages...";
        console.log(chatDiv.value, chatContainer.children[0].value, "prev");
        chatDiv.value = message.receiver; //set chatdiv - msg receiver
        console.log(chatDiv.value, chatContainer.children[0].value);

        chatContainer.style.display = "block";
        chatDiv.innerHTML = "";
        //here blyat)

        chatContainer.children["sendBtnId"].onclick = sendMessage.bind(
          this,
          message.receiver,
          authorId,
          message.author
        );
        break;
      case "lastmessage":
        //append last message in chatbox
        //set in activeChat last message
        if (message.message.sender == chatDiv.value) {
          //activeChat -> if uuid equal == chatValue
          appendLastMessageInActiveChat(
            "receive",
            message.message.sendername,
            message.message.senttime,
            message.message.content
          );
          //not focus
        }

        //update list user lastmessage
        updateDataInListUser(
          message.message.sender,
          message.message.senttime,
          message.message.sendername,
          message.message.content
        );

        //count fix
        el == null
          ? (el = document.getElementById(message.message.sender))
          : null;

        //check if user now - active chat - else count++ & show data

        //update value, else
        if (message.message.sender != chatDiv.value) {
          if (el.childElementCount == 4 && el.children[3].textContent != "") {
            el.children[3].textContent =
              1 + parseInt(el.children[3].textContent);
            el.children[3].classList.add("unread");
          } else {
            let span = document.createElement("span");
            span.id = "unread";
            span.classList.add("unread");
            span.textContent = 1;
            el.append(span);
          }
        }
        //update focused user in chat, first elem in list
        toggleOnlineUser(message.message.sender, "prepend");
        break;
      case "leaveuser":
        //get elem by uuid, set id - beacuse user left
        el == null ? (el = document.getElementById(message.user.uuid)) : null;
        el != null ? el.classList.remove("online") : null;
        el.id = message.user.id; //set id, replace - prev uuid
        console.log(message, "leave");
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
