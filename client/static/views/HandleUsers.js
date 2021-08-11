import { wsConn, getCookie, listMessages, chatStore } from "./WebSocket.js";

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

export const showListUser = (users) => {
  let count = 0;
  if (window.location.pathname == "/chat") {
    let senderUuid = getCookie("session");
    let parent = document.getElementById("userlistbox");

    if (users != null && parent != null) {
      // parent.innerHTML = "";
      for (let [keyUser, user] of Object.entries(users)) {
        let li = document.createElement("li");
        if (Object.entries(users).length == 1) {
          alert("Now, no has online user");
          return;
        }
        user.uuid !== "" ? (li.id = user.uuid) : (li.id = user.id);
        user.online ? ((li.className = "online"), (count += 1)) : "";

        if (user.fullname) {
          let pattern = "";
          // <span class="partner"> ${user.fullname}</span>
          let username = `<h3 class="partner">${user.fullname}</h3>`;
          user.lastmessage["String"] == ""
            ? (pattern = `  ${username}<span> No have messages.. </span>`)
            : (pattern = `${username}
                <span>${user.lastmessage["String"]}</span>
              <span class="time">${user.senttime}</span>
             `);

          li.innerHTML = pattern;

          li.onclick = (e) => {
            toggleOnlineUser(li.id);
            let obj = {
              receiver: li.id,
              sender: senderUuid,
              type: "last10msg",
              offset: 0,
            };
            wsConn.send(JSON.stringify(obj));
            //1 click->get last 10 msg
            //next time - evenListener Scroll()
            //set global array empty, next chatWindwos, own messages
            listMessages.length = 0;
            chatStore.countNewMessage = 0;
          };
        }
        //append without yourself
        console.log(senderUuid, user.uuid);
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
