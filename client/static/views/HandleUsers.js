import { wsConn, getCookie } from "./WebSocket.js";
// export const listUsers = new Map();

export const showListUser = (users) => {
  if (window.location.pathname == "/chat") {
    let senderUuid = getCookie("session");
    let parent = document.getElementById("userlistbox");
    let ul = document.getElementById("listusersID");

    if (users != null && ul != null && parent != null) {
      ul.innerHTML = "";
      //   listUsers.set(user.UUID, user);
      for (let [keyUser, user] of Object.entries(users)) {
        let li = document.createElement("li");
        for (let [key, value] of Object.entries(user)) {
          if (Object.entries(users).length == 1) {
            alert("Now, no has online user");
            return;
          }

          li.id = user.id;

          if (user.online) {
            li.className = "online";
          }

          let uuid = user.uuid;

          if (key == "fullname" && value != "") {
            li.textContent = value;
            li.onclick = (e) => {
              //remove prev clicked elem class
              for (let i = 0; i < ul.children.length; i++) {
                if (ul.children[i].className == "current") {
                  ul.children[i].classList.remove("current");
                }
              }
              li.classList.add("current");

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
          if (key == "uuid") {
            if (value != senderUuid) {
              ul.append(li);
            }
          }
        }
        parent.append(ul);
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
