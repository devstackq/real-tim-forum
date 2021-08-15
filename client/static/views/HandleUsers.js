import { wsConn, getCookie, listMessages, chatStore } from "./WebSocket.js";

export const updateDataInListUser = (receiver, time, sendername, content) => {
  let currentUser = document.getElementById(receiver);
  currentUser != null
    ? ((currentUser.children[2].textContent = time),
      (currentUser.children[1].textContent = `${sendername} : ${content}`))
    : null;
};

//send uuid or id if offline - set/remove current class
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

export const showListUser = (users, wsMessage, type) => {
  if (window.location.pathname == "/chat") {
    if (users != null && parent != null) {
      let senderUuid = getCookie("session");
      let parent = document.getElementById("userlistbox");
      let chatDiv =
        document.querySelector("#message_container").children["chatbox"];
      // if (type == "newuser") {
      //   console.log(users, "new user case,", wsMessage);
      // } else {

      // }
      for (let [keyUser, user] of Object.entries(users)) {
        // console.log("inside ", user);
        let li = document.createElement("li");
        if (Object.entries(users).length == 1) {
          alert("Now, no has online user");
          return;
        }
        //set element id = id || uuid
        user.uuid !== "" ? (li.id = user.uuid) : (li.id = user.id);
        user.online ? (li.className = "online") : "";

        if (user.fullname) {
          let pattern = "";
          let username = `<h3 class="partner">${user.fullname}</h3>`;
          user.lastmessage["String"] == ""
            ? (pattern = `${username}<span> No have messages.. </span>
            <span class="time"></span>
            `)
            : (pattern = `${username}
            <span> ${user.lastsender["String"]} : ${
                user.lastmessage["String"]
              }</span>
              <span class="time">${user.senttime}</span>
              ${
                user.countunread > 0
                  ? ` <span id="unread" class="unread">${user.countunread} </span>`
                  : ""
              }  `);
          li.innerHTML = pattern;

          li.onclick = (e) => {
            //case - newuser
            // set each user active chat.value = uuid || id
            // let chatDiv =
            //   document.querySelector("#message_container").children["chatbox"];
            // chatDiv.value = li.id;
            console.log(li.id, "user click val");

            let obj = {
              sender: senderUuid,
              receiver: li.id,
              type: "last10msg",
              offset: 0,
            };
            wsConn.send(JSON.stringify(obj)); //1 click->get last 10 msg
            toggleOnlineUser(li.id);
            //set global array empty, next chatWindwos, own messages
            listMessages.length = 0;

            //dry
            let current = document.getElementById(li.id);
            //remove unread class &  countUnread = 0
            current != undefined &&
            current.children.length > 3 &&
            current.children[3].id == "unread"
              ? (current.children[3].classList.remove("unread"),
                (current.children[3].textContent = "0"))
              : null;
            // set each user active chat.value = uuid || id
            chatDiv.value = li.id;
          };
        }
        //append without yourself
        if (user.uuid) {
          if (user.uuid != senderUuid) {
            parent.append(li);
          }
        } else if (user.id && !user.uuid) {
          parent.append(li);
        }
      }
    } else {
      alert("no has online user");
    }
  }
};
const openWs = (uuid, type) => {
  wsConn.onopen = () => {
    wsConn.send(
      JSON.stringify({
        sender: uuid,
        type: type,
      })
    );
  };
};
const sendWs = (uuid, type) => {
  wsConn.send(
    JSON.stringify({
      sender: uuid,
      type: type,
    })
  );
};

//mail. pwd, 1 auth,1 open
export const listUsers = (uuid, type) => {
  if (type == "signin" || type == "getusers" || type == "newuser") {
    if (uuid == undefined) {
      uuid = getCookie("session");
    }
    console.log(type, "listuser func", uuid);
    if (getAuthState() == "true") {
      if (wsConn.readyState != 1) {
        openWs(uuid, type);
      } else {
        sendWs(uuid, type);
      }
    }
  }
};
export const getAuthState = () => {
  if (document.cookie.split(";").length > 1) {
    return localStorage.getItem("isAuth");
  }
};
