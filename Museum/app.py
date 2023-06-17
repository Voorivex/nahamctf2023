from flask import Flask, request, render_template, send_from_directory, send_file, redirect, url_for
import os
import urllib
import urllib.request

app = Flask(__name__)

@app.route('/')
def index():
    artifacts = os.listdir(os.path.join(os.getcwd(), 'public'))
    return render_template('index.html', artifacts=artifacts)

@app.route("/public/<file_name>")
def public_sendfile(file_name):
    file_path = os.path.join(os.getcwd(), "public", file_name)
    if not os.path.isfile(file_path):
        return "Error retrieving file", 404
    return send_file(file_path)

@app.route('/browse', methods=['GET'])
def browse():
    file_name = request.args.get('artifact')

    if not file_name:
        return "Please specify the artifact to view.", 400

    artifact_error = "<h1>Artifact not found.</h1>"

    if ".." in file_name:
        return artifact_error, 404

    if file_name[0] == '/' and file_name[1].isalpha():
        return artifact_error, 404
    
    file_path = os.path.join(os.getcwd(), "public", file_name)
    if not os.path.isfile(file_path):
        return artifact_error, 404

    if 'flag.txt' in file_path:
        return "Sorry, sensitive artifacts are not made visible to the public!", 404

    with open(file_path, 'rb') as f:
        data = f.read()

    image_types = ['jpg', 'png', 'gif', 'jpeg']
    if any(file_name.lower().endswith("." + image_type) for image_type in image_types):
        is_image = True
    else:
        is_image = False

    return render_template('view.html', data=data, filename=file_name, is_image=is_image)

@app.route('/submit')
def submit():
    return render_template('submit.html')

@app.route('/private_submission_fetch', methods=['GET'])
def private_submission_fetch():
    url = request.args.get('url')

    if not url:
        return "URL is required.", 400

    response = submission_fetch(url)
    return response

def submission_fetch(url, filename=None):
    return urllib.request.urlretrieve(url, filename=filename)

@app.route('/private_submission')
def private_submission():
    if request.remote_addr != '127.0.0.1':
        return redirect(url_for('submit'))

    url = request.args.get('url')
    file_name = request.args.get('filename')

    if not url or not file_name:
        return "Please specify a URL and a file name.", 400

    try:
        submission_fetch(url, os.path.join(os.getcwd(), 'public', file_name))
    except Exception as e:
        return str(e), 500

    return "Submission received.", 200

if __name__ == '__main__':
    app.run(debug=False, host="0.0.0.0", port=5000)