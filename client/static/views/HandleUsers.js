import { wsConn } from "./WebSocket.js";

export const listUsers = new Map();

// show join, leave - realtime update users
// session update -> user prev close conn user - delete Backend
// if server restart -> all user reconnect

export const showListUser = (users) => {
  let senderUuid = "";
  if (document.cookie.split(";").length > 1) {
    senderUuid = document.cookie.split(";")[0].slice(8).toString();
  }
  console.log(Object.entries(users), 1);
  let parent = document.getElementById("userlistbox");
  let ul = document.getElementById("listusersID");
  ul.innerHTML = "";

  if (users != null && ul != null && parent != null) {
    //   listUsers.set(user.UUID, user);
    for (let [uuid, user] of Object.entries(users)) {
      let li = document.createElement("li");
      for (let [key, value] of Object.entries(user)) {
        if (Object.entries(users).length == 1) {
          // super.showNotify("Now, no has online user", "error");
          alert("Now, no has online user");
          return;
        }
        if (key == "fullname" && value != "") {
          li.textContent = value;
          li.onclick = (e) => {
            //remove prev clicked elem class
            for (let i = 0; i < ul.children.length; i++) {
              if (ul.children[i].className == "current") {
                ul.children[i].classList.remove("current");
              }
            }
            li.className = "current";
            let obj = {
              receiver: uuid,
              sender: senderUuid,
              type: "getmessages",
            };
            wsConn.send(JSON.stringify(obj));
          };
        }
        //append without yourself
        if (uuid != senderUuid) {
          ul.append(li);
        }
      }
      parent.append(ul);
    }
  } else {
    alert("no has online user");
  }
  console.log(listUsers, ":list гыукы");
};

export const addNewUser = (uuid) => {
  console.log("new user1");
  if (getAuthState() == "true") {
    console.log("new user", uuid, wsConn);
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
