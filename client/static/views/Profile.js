import WebSocket from "./WebSocket.js";

export default class Profile extends WebSocket {
  constructor(params) {
    super();
    this.user={added :false}
    this.params = params;
  }
  setTitle(title) {
    document.title = title;
  }

  async init() {
    let response = await fetch("http://localhost:6969/api/profile");
    // console.log(response, "porifle");
    if (response.status === 200) {
      let result = await response.json();
      super.renderSequence(result);
// 1 time call ?
r&d -> ws class in js -> or use -> helper func 

console.log(this.user.added)
if(!this.user.added) {
      let ws = new WebSocket('ws://localhost:6969/api/chat')
      let newuser = {
        sender: super.getUserSession(),
        type: "newuser"
      };
      console.log('send req wss', newuser)
      ws.onopen = () => {
        ws.send(JSON.stringify(newuser));
      };
      this.user.addded = true
      // wtf ?
    }
      // super.openNewWs() //new conn
      // super.addNewUser() //add new user in system
    } else {
      super.showNotify(response.statusText, "error");
      super.showHeader();
      window.location.replace("/signin");
    }
    document.querySelector("#editBio").onclick = () => {
      console.log("edit");
      // let response = await fetch('http://localhost:6969/api/profile/edit')
    };
  }

  async getHtml() {
    let body = `
    <div class="bioUser">
        <button id="editBio"> edit </button>
    </div> 
    <div class="postsUser"></div>
  <div class="votedPost"></div>    `;
    let header = super.showHeader();
    return header + body;
  }
}
