import Parent from "./Parent.js";

export default class Signin extends Parent {
  constructor(text, type, params) {
    super(text, type);
    this.params = params;

    // this.ws = super.getWebsocket();
  }

  setTitle(title) {
    document.title = title;
  }

  async signin() {
    let ok = false;
    let user = {
      email: "",
      password: "",
    };
    user = super.fillObject(user);

    let result = await super.fetch("signin", user);
    if (result !== null) {
      localStorage.setItem("isAuth", true);
      //add user in chat system
      //only send userid in server add new client online
      //input name, message current user
      ok = true;
    } else {
      localStorage.setItem("isAuth", false);
      super.showNotify("incorrect login or password", "error");
    }
    if (ok) {
      super.setOpenWebscoket();
      let ws = super.getWebsocket();

      let newuser = {
        sender: super.getUserSession(),
        type: "newuser",
      };
      //client 1 enter chat service ->
      ws.open = () => {
        ws.send(JSON.stringify(newuser));
      };
      console.log(22, ws.readyState);
      // setTimeout(() => {
      //   window.location.replace("/profile");
      // }, 1500);
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
