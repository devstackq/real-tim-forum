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
        console.log(response, 'porifle')
        if (response.status === 200) {
            let result = await response.json();
            this.showBio(result);
        } else {
            super.showNotify(response.statusText, "error");
            // console.log('not uuid || incorrect')
            // window.location.replace("/signin");
        }
        document.querySelector("#editBio").onclick = () => {
            console.log("edit");
            // let response = await fetch('http://localhost:6969/api/profile/edit')
        };
    }

    async getHtml() {
        let body = `
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
        return header + body
    }
}