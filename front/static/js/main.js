const msgBoxElem = document.getElementById("msg-box")
const backendPath = "/chatroom"
let myName = ""
let connection;

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
const removeOlderBubbles = () => {
	if(msgBoxElem.childElementCount > 50){
		msgBoxElem.removeChild(msgBoxElem.firstChild)
	}
}

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
	connection = new WebSocket("ws://" + backendURL + backendPath)
	connection.onclose = () => {
		console.error("Connection to server lost. Trying to reconnect...")
		connection = new WebSocket("ws://" + backendURL + backendPath)
	}
	connection.onmessage = (e) => {
		const data = JSON.parse(e.data)

		if(data.user === myName){
			data.user = null
		}
		createBubble(data.user, data.msg, getTime(data.time))
		removeOlderBubbles()
	}
	connection.onopen = () => {
		console.log("Connection opened.")
	}

	// Get name from query param
	const name = (new URL(document.location)).searchParams.get("name");
	setName(name)
}