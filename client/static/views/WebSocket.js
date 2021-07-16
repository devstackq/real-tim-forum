import { showListUser, addNewUser } from "./HandleUsers.js";
import { showListMessages } from "./Chat.js";

export let wsConn = null;

export const getSession = () => {
  // console.log(document.cookie.split(";").length, " len");
  if (document.cookie.split(";").length === 3) {
    return document.cookie.split(";")[0].slice(8).toString();
  }
};

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
  wsConn.onmessage = (e) => {
    console.log("getuser sess");
    let message = JSON.parse(e.data);
    switch (message.type) {
      case "observeusers":
        //update user list -> all conns
        console.log("for all users", message.users);
        showListUser(message.users);
        break;
      case "getusers":
        //get user own client
        showListUser(message.users);
        break;
      case "listmessages":
        document.getElementById("notify").value = "";
        // this.HtmlElems.messageContainer.children["chatbox"].style.display =
        //   "block";
        showListMessages(
          message.messages,
          getUserId(),
          getSession(),
          message.users
        );
        break;
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
        console.log(message, "leave user");
        showListUser(message.users);
        //delete here user from listUsers[uuid], rerender
        break;
    }

    wsConn.onclose = function (event) {
      console.log(
        " Обрыв соединения, Код: " + event.code + " причина: " + event.reason
      );
      //   wsConn.send(JSON.stringify({ type: "leave", sender: getSession() }));
      wsConn.close();
      // wsConn.send(JSON.stringify({ type: "close" }));
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
