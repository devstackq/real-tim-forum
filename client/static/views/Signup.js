import Parent from "./Parent.js";

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
    let e = document.getElementById("email").value;
    let p = document.getElementById("password").value;
    let u = document.getElementById("username").value;
    let f = document.getElementById("fName").value;
    let a = document.getElementById("age").value;
    let c = document.getElementById("city").value;
    let g = document.getElementById("gender").value;

    let user = {
      email: e,
      password: p,
      username: u,
      fullname: f,
      age: a,
      city: c,
      gender: g,
    };

    let result = await super.fetch("signup", user);

    if (result > 0 && result != undefined) {
      window.location.replace("http://localhost:6969/signin");
    } else {
      //validParams() todo
      super.showNotify(response.statusText, "error");
      console.log(response.statusText, "error signup");
    }
  }

  init() {
    document.getElementById("signup").onclick = this.signup;
  }

  async getHtml() {
    let body = `
        <div>
        <input type="text" id='fName' required="true" placeholder='full name'>
        <input type='email' id='email' required placeholder='email'>
        <input type="text" id='username' required placeholder='nick'>
        <input type="password" id="password" required placeholder='password'>
        <input type="number" id='age' required placeholder='age'>
        <label> gender
        <select id='gender' placeholder='gender'>
        <option></option>
        <option>man</option>
        <option>woman</option>
      </select>
      </label>
        <input type="text" id='city' required placeholder='city'>
        <input type='submit' id='signup' value="register"/>
        </div>
        `;
    let header = super.showHeader("free");
    return header + body
  }
}
