import { wsConn, getCookie, toggleOnlineUser } from "./WebSocket.js";

export const showListUser = (users) => {
  if (window.location.pathname == "/chat") {
    let senderUuid = getCookie("session");
    let parent = document.getElementById("userlistbox");

    if (users != null && parent != null) {
      // parent.innerHTML = "";
      for (let [keyUser, user] of Object.entries(users)) {
        let li = document.createElement("li");
        // for (let [key, value] of Object.entries(user)) {
        if (Object.entries(users).length == 1) {
          alert("Now, no has online user");
          return;
        }
        user.uuid == "" ? (li.id = user.id) : (li.id = user.uuid);

        user.online ? (li.className = "online") : "";

        let uuid = user.uuid;
        console.log(keyUser, user, Object.entries(user.lastmessage)[0]);

        // if (key == "fullname" && value != "") {
        if (user.fullname) {
          li.textContent = `Name: ${user.fullname}  Last Message:${user.lastmessage["String"]}  From:${user.lastsender["String"]}  Time:${user.senttime["Time"]} `;
          li.onclick = (e) => {
            //remove prev clicked elem class, //dry /
            toggleOnlineUser(li.id);

            if (uuid == "") {
              uuid = user.id.toString();
            }
            let obj = {
              receiver: uuid,
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
        // parent.append(ul);
      }
    } else {
      alert("no has online user");
    }
  }
};
//mail. pwd, 1 auth,1 open
export const addNewUser = (uuid) => {
  if (uuid == undefined) {
    uuid = getCookie("session");
  }
  if (getAuthState() == "true") {
    // console.log("new user", uuid, wsConn);
    wsConn.onopen = () =>
      wsConn.send(
        JSON.stringify({
          sender: uuid,
          type: "newuser",
        })
      );
  }
};

export const getAuthState = () => {
  if (document.cookie.split(";").length > 1) {
    return localStorage.getItem("isAuth");
  }
};
