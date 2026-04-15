<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Interactive Real-Time Dashboard</title>
    <link rel="stylesheet" href="styles.css">
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f4f4f4;
            padding: 20px;
        }
        .dashboard {
            display: flex;
            flex-direction: column;
            align-items: center;
        }
        .card {
            background: white;
            border-radius: 8px;
            box-shadow: 0 2px 5px rgba(0,0,0,0.1);
            margin: 10px;
            padding: 20px;
            width: 300px;
            text-align: center;
            position: relative;
            overflow: hidden;
            transition: transform 0.3s;
        }
        .card:hover {
            transform: scale(1.05);
        }
        .confidence-score,
        .activity-level {
            font-size: 1.5em;
            margin-top: 10px;
        }
        #wsConnectionStatus {
            font-size: 1.2em;
            color: red;
        }
    </style>
</head>
<body>
    <div class="dashboard">
        <h1>Live Prediction Dashboard</h1>
        <div id="wsConnectionStatus">Connecting...</div>
        <div class="card">
            <h2>Live Prediction</h2>
            <div id="prediction">Loading...</div>
            <div class="confidence-score">Confidence: <span id="confidence">-</span>%</div>
            <div class="activity-level">Activity Level: <span id="activity">-</span></div>
        </div>
        <div class="card">
            <h2>Current Status</h2>
            <div id="status">Loading status...</div>
        </div>
    </div>
    <script>
        const ws = new WebSocket('ws://your-websocket-url');
        const predictionElement = document.getElementById('prediction');
        const confidenceElement = document.getElementById('confidence');
        const activityElement = document.getElementById('activity');
        const statusElement = document.getElementById('status');
        const connectionStatusElement = document.getElementById('wsConnectionStatus');

        ws.onopen = function() {
            connectionStatusElement.innerText = 'Connected!';
            connectionStatusElement.style.color = 'green';
        }; 

        ws.onmessage = function(event) {
            const data = JSON.parse(event.data);
            predictionElement.innerText = data.prediction;
            confidenceElement.innerText = data.confidence;
            activityElement.innerText = data.activity_level;
            statusElement.innerText = data.status;
        };

        ws.onclose = function() {
            connectionStatusElement.innerText = 'Disconnected';
            connectionStatusElement.style.color = 'red';
        };
    </script>
</body>
</html>