import { wsConn, getCookie, toggleOnlineUser } from "./WebSocket.js";

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
          user.lastmessage["String"] == ""
            ? (pattern = `No have messages with:  ${user.fullname}`)
            : (pattern = `<h3 class="partner">${user.fullname}</h3>
                <span>${user.lastmessage["String"]}</span>
              <span class="time">${user.senttime["Time"]}</span>
             `);

          li.innerHTML = pattern;
          // li.append(internlocutor);
          li.onclick = (e) => {
            //remove prev clicked elem class, //dry /
            toggleOnlineUser(li.id);
            let obj = {
              receiver: li.id,
              sender: senderUuid,
              type: "getmessages",
            };
            wsConn.send(JSON.stringify(obj));
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
