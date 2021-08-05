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
        // for (let [key, value] of Object.entries(user)) {
        if (Object.entries(users).length == 1) {
          // console.log(users.length, Object.entries(users).length);
          alert("Now, no has online user");
          return;
        }
        user.uuid == "" ? (li.id = user.id) : (li.id = user.uuid);
        user.online ? ((li.className = "online"), (count += 1)) : "";

        let uuid = user.uuid;

        if (user.fullname) {
          let pattern = "";
          let internlocutor = document.createElement("span");
          internlocutor.className = "internlocutor";
          internlocutor.textContent = user.fullname;
          user.lastmessage["String"] == ""
            ? (pattern = `Now now have messages with:  `)
            : (pattern = `From : ${user.lastsender["String"]} \n Message: ${user.lastmessage["String"]} \n Time:${user.senttime["Time"]} Chat with: `);
          li.textContent = pattern;
          // li.textContent = u
          li.append(internlocutor);
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
const sendWs = (uuid, type)=> {
  wsConn.send(
    JSON.stringify({
      sender: uuid,
      type: type,
    })
  );
};

//mail. pwd, 1 auth,1 open
export const listUsers = (uuid, type) => {

  if (type == "signin" || type == "getusers") {
    if (uuid == undefined) {
      uuid = getCookie("session");
    }
    if (getAuthState() == "true") {
      if (wsConn.readyState != 1) {
        openWs(uuid, "online");
      } else {
      sendWs(uuid, "online")
      }
  }
}
}
export const getAuthState = () => {
  if (document.cookie.split(";").length > 1) {
    return localStorage.getItem("isAuth");
  }
};
