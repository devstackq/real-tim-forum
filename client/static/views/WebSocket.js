import { showListUser, addNewUser } from "./HandleUsers.js";

export let wsConn = null;
export let uuid = null;

export const wsInit = (type, ...args) => {
  if (uuid == null) {
    uuid = args[0];
    //set uuid
    if (uuid == null) {
      if (document.cookie.split(";").length > 1) {
        uuid = document.cookie.split(";")[0].slice(8).toString();
      }
    }
  }

  if (wsConn == null) {
    console.log(wsConn, "val, singleton?");
    wsConn = new WebSocket("ws://localhost:6969/api/chat");
  }
  if (type == "signin" && wsConn != null) {
    addNewUser(args[0]);
  } else if (type == "chat" && wsConn != null) {
    if (wsConn != null) {
      if (wsConn.readyState == 1) {
        wsConn.send(JSON.stringify({ type: "getusers" }));
      }
    }
  }

  // if (super.getAuthState() == "true") {
  wsConn.onmessage = (e) => {
    let message = JSON.parse(e.data);
    // let time = new Date(message.date);
    switch (message.type) {
      case "observeusers":
        if (type == "chat") {
          //show Uniq users Map /
          console.log(" for all users", message.users);
          showListUser(message.users, uuid);
        }
        break;
      case "getusers":
        if (type == "chat") {
          console.log("get users", uuid);
          showListUser(message.users, uuid);
        }
        break;
      // case "listmessages":
      //   document.getElementById("notify").value = "";
      //   this.HtmlElems.messageContainer.children["chatbox"].style.display =
      //     "block";
      //   this.showListMessage(message.messages);
      //   break;
      // case "nomessages":
      //   this.HtmlElems.messageContainer.children["chatbox"].style.display =
      //     "none";
      //   this.HtmlElems.messageContainer.style.display = "block";
      //   this.HtmlElems.messageContainer.children["chatbox"].innerHTML = "";
      //   //now no messages -> fix, show message field
      //   this.showChatWindow(this, message.receiver);
      //   super.showNotify("now no messages", "error");
      //   break;

      // case "lastmessage":
      //   let span = document.createElement("span");
      //   let div = document.createElement("div");
      //   //dry
      //   // fix receive - sender -> message -> correct show name
      //   //chat setInterval work only if -> /router -> caht
      //   //           left join - refactor
      //   // fix - another user choice 1 ->hide / show
      //   //how much time call create db ?
      //   span.textContent = `${message.message.aname} : \n ${message.message.content} ${message.message.senttime}  `;
      //   div.append(span);
      //   this.HtmlElems.messageContainer.children["chatbox"].append(div);
      //   this.HtmlElems.messageContainer.children["messageFieldId"].value = "";
      //   break;
      case "leave":
        // this.onlineUsers.delete(message.receiver);
        // console.log(this.onlineUsers);
        console.log(message, "leave user");
        break;
    }

    wsConn.onclose = function (event) {
      console.log("Обрыв соединения");
      console.log("Код: " + event.code + " причина: " + event.reason);
    };
    wsConn.onerror = function (error) {
      console.log("Ошибка " + error.message);
    };
  };
};
// ws.js
//import chathandler from ..
// export let conn;
// wsStart(){
//   conn = new WS()
//   onmsessage() {
// case 1://
//  chathandler.ShowUser(data)
// case 2 :
// chathandler.getMessage(data)
//   }
// }

// index.js
// wsStart()

// chat.js
// new Chat(user, conn)

// class {

// }
