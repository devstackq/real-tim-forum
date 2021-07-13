import { wsConn } from "./WebSocket.js";

export const listUsers = new Map();

show correct user, 
show join, leave - realtime update users

export const showListUser = (users, uuid) => {
  let parent = document.getElementById("userlistbox");
  let ul = document.getElementById("listusersID");
  ul.innerHTML = "";
  // iter obj users
  if (users != null) {
    for (let [k, user] of Object.entries(users)) {
      listUsers.set(user.UUID, user);
      // listUsers.add(user);
    }

    for (let [k, user] of Object.entries(Array.from(listUsers))) {
      // this.onlineUsers.set(key, user);
      let li = "";
      let uzik = Object.entries(user[1]);
      // let objUuid = Object.entries(user[0]);
      for (let [key, value] of uzik) {
        if (uzik.length == 1) {
          // super.showNotify("Now, no has online user", "error");
          alert("Now, no has online user");
          return;
        }
        if (key == "fullname") {
          li = document.createElement("li");
          li.textContent = value;
          li.onclick = (e) => {
            //remove prev clicked elem class
            for (let i = 0; i < ul.children.length; i++) {
              if (ul.children[i].className == "current") {
                ul.children[i].classList.remove("current");
              }
            }
            li.className = "current";
            // this.HtmlElems.messageContainer.children["chatbox"].style.display =
            //   "none";
            // chatbox.innerHTML = "";
            let obj = {
              receiver: uzik.UUID,
              sender: uuid,
              type: "getmessages",
            };
            wsConn.send(JSON.stringify(obj));
          };
        }
        if (uzik.UUID != uuid && uzik.length > 1) {
          ul.append(li);
        }
      }
    }
    parent.append(ul);

    console.log(listUsers, "list", ul.children);
  } else {
    alert("no has online user");
  }
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
