export default class Parent {

    constructor(text, type) {
        this.text = text
        this.type = type
        this.value = ''
    }
    getPostById(id) {
        console.log(id, 'id')
    }


    render(item, idx, where) {

        let wrapper = document.querySelector(where)
        let btn = document.createElement('button')

        for (let [k, v] of Object.entries(item)) {

            if (v != "" && v != null) {
                let span = document.createElement('span')
                span.id = k
                span.textContent = ` ${k} : ${v}`
                    // btn.id = 'getPostBy' + idx
                btn.onclick = async function() {
                    // this.getPostById(idx)
                    let postId = { value: 0 }
                    postId.value = idx
                    let response = await fetch('http://localhost:6969/api/post/id', {
                        method: 'POST',
                        headers: {
                            'Content-Type': 'application/json;charset=utf-8'
                        },
                        body: JSON.stringify(postId)
                    });
                    if (response.status == 200) {
                        console.log(response.statusText)
                            //redirect post component
                            // or here render () ?
                    }
                };
                btn.value = idx
                btn.textContent = `see post`
                wrapper.appendChild(span)
            }
            wrapper.appendChild(btn)
        }
        //send backend - postId
        // let btn = document.querySelector("#getPostById")
        // console.log(btn)
        // btn.onclick = this.getPostById(btn.value)
    }


    showHeader(type) {

        let login = ""
        let register = ""
        let logout = ""
        let profile = ""

        if (type == 'free') {
            profile = ""
            logout = ""
            register = `<a href="/signup"  class="nav__link signup" data-link>Signup</a>`
            login = `<a href="/signin"  class="nav__link signin" data-link>Signin</a>`
        } else if (type == 'auth') {
            register = ""
            login = ""
            logout = `<a href="/logout" id='logout' class="nav__link logout" data-link>Logout</a>`
            profile = `<a href="/profile" class="nav__link" data-link>Profile</a>`
        }

        return `
        <nav class="nav">
        <a href="/all" class="nav__link" data-link>Main</a>
        ${login}
        ${register}
        <div class="dropdown">
          <button class="dropbtn">Categories</button>
          <div class="dropdown-content">
          <a href="/love" data-link>love</a>
          <a href="/science" data-link>science</a>
          <a href="/nature" data-link>nature</a>
        </div>
        </div>
       ${profile}
        ${logout}
    </nav>
    <span class='notify' > </span> 
`
    }

    showNotify(text, type) {

        let notify = document.getElementsByClassName('notify')[0]

        if (type == 'error') {
            notify.style.display = 'block'
            notify.textContent = text
        } else if (type == 'hide') {
            notify.style.display = 'none'
        }
    }
}