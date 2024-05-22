const data = {
    "method": "ComputeService.Multiply",
    "params": [
        {
            "a": 125,
            "b": 30
        }
    ],
    "id": 0
}

const options = {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json'
    },
    body: JSON.stringify(data)
}

fetch('http://localhost:8080/compute', options)
    .then(res=> res.json())
    .then(json => {
        console.log(json)
    })

