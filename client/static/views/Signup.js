import Parent from "./Parent.js";
import { redirect } from "../index.js";
import { wsInit } from "./WebSocket.js";

export default class Signup extends Parent {
  constructor(params) {
    super();
    this.params = params;
    this.name = name;
  }
  setTitle(title) {
    document.title = title;
  }

  async signup() {
    let user = {
      email: "",
      password: "",
      username: "",
      fullname: "",
      age: 0,
      city: "",
      gender: "",
    };

    user = super.fillObject(user);

    //success signup user return uid
    let response = await super.fetch("signup", user);
    if (response.status === 200) {
      //success signup
      let result = await super.fetch("signin", {
        email: user.email,
        password: user.password,
      });
      if (result != null) {
        //signin success
        localStorage.setItem("isAuth", true);
        wsInit(result.uuid, "newuser");
        console.log(result.uuid, "signup");
        //newuser type, another user update list user
        redirect("profile");
      } else {
        console.log("no correct login or password");
      }
    } else {
      console.log(response.value);
      let j = await response.json();
      super.showNotify(j.value, "error");
    }
  }
  init() {
    document.getElementById("signup").onclick = this.signup;
  }

  async getHtml() {
    let body = `
        <div class="signup_container">
        <h3> Signup with us! </h3>
        <input type="text" id='fullname' required="true" placeholder='full name'>
        <input type='email' id='email' required placeholder='email'>
        <input type="text" id='username' required placeholder='nick'>
        <input type="password" id="password" required placeholder='password, 123User!'>
        <input type="number" min="14" max="99" id='age' required placeholder='age'>
        <label> gender
        <select id='gender' placeholder='gender'>
        <option></option>
        <option>man</option>
        <option>woman</option>
      </select>
      </label>
        <input type="text" id='city' required placeholder='city'>
        <input type='submit' id='signup' value="Register"/>
        
        <div class='is_register'> <span>if registered go to: </span> 
        <a href='signin' > signin </a>
      </div>
        </div>
        `;
    let header = super.showHeader();
    return header + body;
  }
}
