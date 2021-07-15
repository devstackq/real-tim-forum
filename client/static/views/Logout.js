import Parent from "./Parent.js";
import { wsConn, getSession } from "./WebSocket.js";
import { listUsers } from "./HandleUsers.js";
import router from "../index.js";

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

      wsConn.send(JSON.stringify({ type: "leave", sender: getSession() }));
      // listUsers.delete(uuid);
      document.cookie = "session=; expires=Thu, 01 Jan 1970 00:00:01 GMT;";
      document.cookie = "user_id=; expires=Thu, 01 Jan 1970 00:00:01 GMT;";
      localStorage.setItem("isAuth", false);

      history.pushState(null, "profile", "http://localhost:6969/all");
      window.addEventListener("popstate", router());
      // window.location.replace("/all");
    } else {
      console.log("error logout");
      super.showNotify(response.statusText, "error");
    }
  }
  async getHtml() {
    return super.showHeader();
  }
}
