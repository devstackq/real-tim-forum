import Parent from "./Parent.js";
import { wsConn, getCookie } from "./WebSocket.js";
import { redirect } from "../index.js";

export default class extends Parent {
  constructor(params) {
    super();
    this.params = params;
  }
  setTitle(title) {
    document.title = title;
  }

  async init() {
    let response = await fetch("http://localhost:6969/api/logout");
    if (response.status === 200) {
      //delete cookie & auth state false
      console.log(wsConn.readyState, getCookie("session"))
norm send json

      wsConn.send(
        JSON.stringify({
          type: "leave",
          sender: getCookie("session"),
          userid: parseInt(getCookie("user_id")),
        })
      );

      // listUsers.delete(uuid);
      document.cookie = "session=; expires=Thu, 01 Jan 1970 00:00:01 GMT;";
      document.cookie = "user_id=; expires=Thu, 01 Jan 1970 00:00:01 GMT;";
      localStorage.setItem("isAuth", false);
      redirect("all");
      window.location.reload();
    } else {
      console.log("error logout");
      super.showNotify(response.statusText, "error");
    }
  }
  async getHtml() {
    return super.showHeader();
  }
}
