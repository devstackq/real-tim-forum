import Parent from "./Parent.js";

export default class Profile extends Parent {
  constructor(params) {
    super();
    this.params = params;
  }
  setTitle(title) {
    document.title = title;
  }
  //auto create or create then -> fill data ?
  render(object, where, ...type) {
    // console.log(object, object.User['email'])
    let parent = document.querySelector(where);
    let postsHtml = document.querySelector(".profilePost");

    if(object.User != null) {
      // this.render(object['User'], where)
      for (let i = 0; i < parent.children.length; i++) {
        parent.children[i].textContent = `${parent.children[i].id } : ${object.User[parent.children[i].id]} `
       } 
    }
    if(object.Posts != null) {
      object.Posts.forEach(item => {
      for(let [i , v ] of Object.entries(item)) {
        console.log(i,v)
        dynamic create html dield for each post -> then fill data
        postsHtml.children[4].textContent = `${i} : ${v}`
      }
    })
      
    }
      // if(type.length > 1) {
          // for(let i =0; i < type.length; i++) {
          //     if(object[type[i]].length != undefined) {
          //         this.render(object[type[i]], where)
          //         // console.log( object[type[i]], object[type[i]].length)
          //       }
          // }
        }

  showBio(data) {
    let bio = document.querySelector(".profileBio");
    for (let i = 0; i < bio.children.length; i++) {
      for (let [k, v] of Object.entries(data)) {
        if (k == bio.children[i].id) {
          bio.children[i].textContent = ` ${k} : ${v}`;
        }
      }
    }
  }

  async init() {
    let response = await fetch("http://localhost:6969/api/profile");
    console.log(response, "porifle");
    if (response.status === 200) {
      let result = await response.json();
    // console.log(result.Posts)
    this.render(result, '.profileBio', "Posts", "User")
    //   this.showBio(result);
    } else {
      super.showNotify(response.statusText, "error");
      // console.log('not uuid || incorrect')
      window.location.replace("/signin");
    }
    document.querySelector("#editBio").onclick = () => {
      console.log("edit");
      // let response = await fetch('http://localhost:6969/api/profile/edit')
    };
  }

  async getHtml() {
    let body = `
    <div class="profilePost">
    <p id="thread"> </p>
    <p id="content"> </p>
    <p id="countlike"> </p>
    <p id="countdislike"> </p>
    </div>
    
    <div class="profileBio">
        <p id="fullname"> </p>
        <p id="email"> </p>
        <p id="age"> </p>
        <p id="gender"> </p>
        <p id="city"> </p>
        <p id="username"> </p>
        <p id="lastseen"> </p>
        <button id="editBio"> edit </button>
    </div> 
    `;
    let header = super.showHeader("auth");
    return header + body;
  }
}
