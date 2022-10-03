//setInterval(() => {location.reload()}, 1000)

const msgBoxElem = document.getElementById("msg-box")

const createBubble = (sender, msg, time) => {
	const bubbleElem = document.createElement("div")

	if(sender == null){
		bubbleElem.classList.add("my-bubble")
	} else {
		bubbleElem.classList.add("others-bubble")
		const senderElem = document.createElement("div")
		senderElem.classList.add("sender")
		senderElem.innerHTML = sender
		bubbleElem.appendChild(senderElem)
	}
	const msgElem = document.createElement("div")
	msgElem.innerHTML = msg
	bubbleElem.appendChild(msgElem)

	const blankSpaceElem = document.createElement("span")
	blankSpaceElem.classList.add("blank-space")
	bubbleElem.appendChild(blankSpaceElem)

	const timestampElem = document.createElement("div")
	timestampElem.classList.add("timestamp")
	timestampElem.innerHTML = time
	bubbleElem.appendChild(timestampElem)

	msgBoxElem.appendChild(bubbleElem)

	msgBoxElem.scrollTo(0,msgBoxElem.scrollHeight)
}
const getTime = (t) => {
	const d = new Date(t)
	const hours = d.getHours() > 9 ? d.getHours() : "0"+d.getHours()
	const minutes = d.getMinutes() > 9 ? d.getMinutes() : "0"+d.getMinutes()

	return hours + ":" + minutes
}
const setName = (name) => {
	myName = name
	document.getElementById("name").innerHTML = myName
}

const backendURL = "192.168.0.34:8080/chatroom"
let myName = "Don Pepito"
let connection;

document.getElementById("send-btn").addEventListener("click", () => {
	const msg = document.getElementById("msg-input")

	if(msg.value === "") return

	const data = JSON.stringify({"user": myName, "msg": msg.value})
	connection.send(data)

	msg.value = ""
})
document.getElementById('msg-input').addEventListener('keydown', (e) => {
	if(e.key !== 'Enter') return

	document.getElementById('send-btn').click()
})

window.onload = () => {
	connection = new WebSocket("ws://" + backendURL)
	connection.onclose = () => {
		console.error("Connection to server lost.")
	}
	connection.onmessage = (e) => {
		const data = JSON.parse(e.data)

		if(data.user === myName){
			data.user = null
		}
		createBubble(data.user, data.msg, getTime(data.time))
	}
	connection.onopen = () => {
		console.log("Connection opened.")
	}

	// Get name from query param
	const name = (new URL(document.location)).searchParams.get("name");
	setName(name)
}