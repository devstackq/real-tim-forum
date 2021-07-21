import { wsConn, getCookie } from "./WebSocket.js";

// export const listUsers = new Map();
export let listUsers = {};

export const showListUser = (users) => {
 
  if(users.length > 1 ) {
  listUsers=  users
 }
 console.log(users)

  if (window.location.pathname == "/chat") {
    let senderUuid = "";


    senderUuid = getCookie("session");
    let parent = document.getElementById("userlistbox");
    let ul = document.getElementById("listusersID");

    if (users != null && ul != null && parent != null) {
      ul.innerHTML = "";
      //   listUsers.set(user.UUID, user);
      for (let [uuid, user] of Object.entries(listUsers)) {
        let li = document.createElement("li");
        for (let [key, value] of Object.entries(user)) {
          if (Object.entries(listUsers).length == 1) {
            // super.showNotify("Now, no has online user", "error");
            alert("Now, no has online user");
            return;
          }
          if (key == "fullname" && value != "") {
            //dry
            if (key =="online") {
              if (value) {
                li.className = 'online'
              }
            }else {
              li.className = 'offline'
            }
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
    console.log(listUsers, ":list user in <Map");
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
