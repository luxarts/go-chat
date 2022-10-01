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
}
const getCurrentTime = () => {
	const d = new Date()
	const hours = d.getHours() > 9 ? d.getHours() : "0"+d.getHours()
	const minutes = d.getMinutes() > 9 ? d.getMinutes() : "0"+d.getMinutes()

	return hours + ":" + minutes
}

document.getElementById("send-btn").addEventListener("click", () => {
	const msg = document.getElementById("msg-input")

	if(msg.value === "") return

	const time = getCurrentTime()

	createBubble(null, msg.value, time)

	msg.value = ""
})


// createBubble("Don José", "Hola don Pepito", "23:58")
// createBubble(null, "Hola don José", "23:59")
// createBubble("Don José", "Hola don José", "23:58")
// createBubble(null, "Hola don Pepito", "23:59")