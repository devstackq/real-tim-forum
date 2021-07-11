import Parent from "./Parent.js";

export default class Profile extends Parent {
  constructor(params) {
    super();
    this.params = params;
    this.ws = super.getWebsocket()   

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
      let newuser = {
        sender: super.getUserSession(),
        type: "newuser",
      };
      //client 1 enter chat service ->
      this.ws.onopen = () => this.ws.send(JSON.stringify(newuser));
  
    } else {
      super.showNotify(response.statusText, "error");
      super.showHeader()
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
