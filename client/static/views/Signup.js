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
    // console.log(user,0)
    if (user.city == "") {
      user.city = "Almaty";
    } else if (user.gender == "") {
      user.gender = "man";
    } else if (user.age == 0) {
      user.age = 21;
    }
    //success signup user return uid
    let uid = await super.fetch("signup", user);
    console.log(uid);

    if (uid > 0) {
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
      console.log("error signup");
      //validParams() todo
      super.showNotify("signup error", "error");
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
        <input type="password" id="password" required placeholder='password'>
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
