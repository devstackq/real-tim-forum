import { showListUser, listUsers } from "./HandleUsers.js";
import { showListMessages, sendMessage } from "./Chat.js";

export let ListUsers = {};
export let wsConn = null;
export let authorName = "";
export let listMessages = [];
export let countNewMessage = { value: 0 };

let chatDiv;

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
//send uuid or id if offline
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

const prepareUserButton = (uuid, fullname, idx) => {
  let listUsersDom = document.querySelector("#userlistbox");

  let el = document.createElement("li");
  el.id = uuid;
  el.classList.add("online");
  el.textContent = "Now no have message with:" + fullname;
  listUsersDom.insertBefore(el, listUsersDom.children[idx]);
  el.onclick = (e) => {
    //remove prev clicked elem class, //dry /
    toggleOnlineUser(uuid);
    let obj = {
      receiver: uuid,
      sender: getCookie("session"),
      type: "getmessages",
    };
    wsConn.send(JSON.stringify(obj));
  };
};

const insertNewUser = (message, tempListUsers) => {
  if (tempListUsers.length > 1) {
    for (let [index, user] of Object.entries(tempListUsers)) {
      if (`${user.lastmessage["String"]}` == "") {
        if (message.user.fullname < user.fullname) {
          prepareUserButton(message.user.uuid, message.user.fullname, index);
          break;
        }
      }
    }
  } else {
    prepareUserButton(message.user.uuid, message.user.fullname, 0);
  }
};

//update listMsg, call listMsgs()
const getMessages = (chatDiv, receiver, sender, type, offset) => {
  //list message get, append in array, - then scrollView -> 0-> down, last element
  // if (chatDiv.scrollTop === 0) {
  console.log("send query", "ofset val: ", offset);
  let obj = {
    receiver: receiver,
    sender: sender,
    type: type,
    offset: (offset += 10),
  };
  wsConn.send(JSON.stringify(obj));
  // chatDiv.removeEventListener("scroll", getMessages);
  return;
  // }
};

function throttle(fn, wait) {
  var time = Date.now();
  return function () {
    if (time + wait - Date.now() < 0) {
      console.log("call func");
      fn();
      time = Date.now();
    }
  };
}

function scrollToBottom() {
  chatDiv.scrollTop = chatDiv.scrollHeight;
}

//call acc * 2?
function debounce(func, timeout) {
  let timer;
  return (...args) => {
    clearTimeout(timer);
    timer = setTimeout(() => {
      func.apply(this, args);
    }, timeout);
  };
}
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

  let chatContainer = document.querySelector("#message_container");

  let tempListUsers = [];
  if (chatContainer != null) {
    chatDiv = chatContainer.children["chatbox"];
  }
  let count = false;

  wsConn.onmessage = (e) => {
    let authorId = getCookie("user_id");
    let authorSession = getCookie("session");
    let message = JSON.parse(e.data);
    // listUsers = message.users;
    console.log(message.type);
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
      case "newuser":
        //signup user, another online user -show him
        insertNewUser(message, tempListUsers);
        break;
      case "online":
        el == null ? (el = document.getElementById(message.user.uuid)) : null;
        el != null ? (el.className = "online") : null;
        el.id = message.user.uuid;
        //1 set user onlien state, update uuid
        break;
      case "observeusers":
        //temp, for sort & insert  in DOm, new signup user
        tempListUsers = [];
        tempListUsers = [...message.users];

        authorName = message.author;
        showListUser(message.users);
        break;
      case "listmessages":
        document.getElementById("notify").value = "";
        //append -> current merge 2 array
        listMessages = [...message.messages.reverse(), ...listMessages]; // for compare, & ignoring duplicate msg
        console.log(
          listMessages.length,
          "len msg",
          listMessages,
          "count new msg",
          countNewMessage
        );

        showListMessages(listMessages, authorId, authorSession, message.author);
        // scroll -> up to 10 MSGesture, position -> sned rRequest
        chatDiv.value = message.receiver;
        chatDiv.children[chatDiv.children.length - 1].scrollIntoView();

        //scroll -> throttle, ms -> call after ms, fn
        //get count all msg, throttle called count: ++, if allMsgCount / 10 == countThrottle, not work throttle
        count = false;
        console.log(count, "ccc");

        chatDiv.addEventListener("scroll", () => {
          if (count) {
            console.log(count, "count");
            return;
          } else {
            console.log(chatDiv.scrollTop);
            if (chatDiv.scrollTop == 0) {
              // debounce(
              let obj = {
                receiver: message.receiver,
                sender: message.sender,
                type: "last10msg",
                offset: message.offset + countNewMessage.value,
              };
              wsConn.send(JSON.stringify(obj));
              count = true;

              // getMessages.bind(
              //   this,
              //   chatDiv,
              //   message.receiver,
              //   message.sender,
              //   "last10msg",
              //   message.offset + countNewMessage.value
              // );
              //   300
              // );
            }
          }
        });

        break;
      case "nomessages":
        alert("no have messages..");
        document.getElementById("notify").value = "no have messages...";
        chatContainer.style.display = "block";
        chatDiv.innerHTML = "";

        chatContainer.children["sendBtnId"].onclick = sendMessage.bind(
          this,
          message.receiver,
          authorId,
          message.author
        );
        break;
      case "lastmessage":
        countNewMessage.value += 1;
        //append last m.vaessage, chatbox
        //dry ?
        let div = document.createElement("div");
        let span = document.createElement("span");
        let text = `  ${message.message.sendername} ${message.message.content}  ${message.message.senttime}`;
        span.style.padding = "9px";
        span.textContent = text;
        //active windwos -> if uuid equal
        if (chatDiv.value == message.message.sender) {
          div.append(span);
          chatDiv.append(div);
          chatDiv.children[chatDiv.children.length - 1].scrollIntoView();
        }
        //list users - update messages
        el == null
          ? (el = document.getElementById(message.message.sender))
          : null;
        el.textContent = text;
        chatContainer.children["messageFieldId"].value = "";
        //update focused user in chat
        toggleOnlineUser(message.message.sender, "prepend");
        break;
      case "leave":
        //get elem by uuid, set id - beacuse user left
        el == null ? (el = document.getElementById(message.user.uuid)) : null;
        el != null ? el.classList.remove("online") : null;
        el.id = message.user.id; //set id, replace - prev uuid
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
