import json
import os
import requests

def lambda_handler(event, context):
    sns = event['Records'][0]['Sns']
    push_title = sns['Subject']
    push_body = sns['Message']

    data = json.dumps({'title': push_title, 'body': push_body, 'type': 'note'})

    pushbullet_access_token = os.environ.get('PUSHBULLET_ACCESS_TOKEN')
    
    response = requests.post('https://api.pushbullet.com/v2/pushes', \
        headers={'Access-Token': pushbullet_access_token, 'Content-Type': 'application/json'}, \
        data=data
    )

    if response.status_code == 200:
        print('Successfully sent push')
    else:
        print('Push failed due to HTTP ' + str(response.status_code))

    return {
        'statusCode': response.status_code,
        'body': json.dumps(push_title+': '+push_body)
    }
