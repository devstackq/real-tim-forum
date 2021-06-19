import Parent from "./Parent.js";

export default class extends Parent {
  constructor(params) {
    super();
    this.params = params;
  }
  setTitle(title) {
    document.title = title;
  }

  async logout() {}

   init() {
    document.querySelector("#logout").onclick = async() => {
        let response = await fetch("http://localhost:6969/api/logout");
      if (response.status === 200) {
        //delete cookie & auth state false
        document.cookie = "session=; expires=Thu, 01 Jan 1970 00:00:01 GMT;";
        document.cookie = "user_id=; expires=Thu, 01 Jan 1970 00:00:01 GMT;";
        localStorage.setItem("isAuth", false);
        window.location.replace("/all");
      } else {
        console.log("error logout");
        super.showNotify(response.statusText, "error");
      }
    };
  }

  async getHtml() {
    return super.showHeader("auth");
  }
}
