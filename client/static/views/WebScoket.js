import Parent from "./Parent.js";

export default class WebSocket extends Parent {
  constructor(text, type) {
    super();
    this.ws = {
      conn: new WebSocket("ws://localhost:6969/api/chat"),
    };
  }

  getWs() {
    return this.ws.conn;
  }
}
