import Parent from "./Parent.js";

export default class Signin extends Parent {
  constructor(text, type, params) {
    super(text, type);
    this.params = params;
  }

  setTitle(title) {
    document.title = title;
  }

  async signin() {
    let user = {
      email: "",
      password: "",
    };
    user = super.fillObject(user);

    let result = await super.fetch("signin", user);
    if (result !== null) {
      localStorage.setItem("isAuth", true);
   //new ws - addnewuser -> work 2 ws - chat.js ?
      window.location.replace("/profile");
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
        <div>
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
