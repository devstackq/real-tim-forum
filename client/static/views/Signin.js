import Parent from "./Parent.js";
import { redirect } from "../index.js";
import { wsInit } from "./WebSocket.js";

export default class Signin extends Parent {
  constructor(params) {
    super();
    this.params = params;
  }

  setTitle(title) {
    document.title = title;
  }

  async signin() {
    let user = {
      username: "",
      email: "",
      password: "",
    };
    //prepare object - see field
    user = super.fillObject(user);
    if (user.email.match(/@/g)) {
      user.username = "";
    } else {
      user.username = user.email;
      user.email = "";
    }

    let result = await super.fetch("signin", user);
    if (result.status != 400) {
      localStorage.setItem("isAuth", true);
      wsInit(result.uuid, "signin");
      redirect("profile");
    } else {
      localStorage.setItem("isAuth", false);
      super.showNotify("incorrect login or password", "error");
    }
  }

  init() {
    localStorage.setItem("isAuth", false);
    document.querySelector("#signin").onclick = this.signin.bind(this);
  }

  async getHtml() {
    let body = `
        <div class="signin_wrapper">
        <h3> Signin</h3>
        <input type='email' id='email' placeholder='email' required>
        <input type="password" id="password" placeholder='password' required>
        <input type='submit' id='signin' value="signin"/>
        <div> <span>if not register go to: </span> 
         <a href='signup' > signup </a>
       </div>
         </div>
        `;
    let header = super.showHeader();
    return header + body;
  }
}
