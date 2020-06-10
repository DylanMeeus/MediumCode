import cv2
import io
from PIL import Image
import numpy as np
import sys 
import os
import base64
import json
import requests

dploy_url = "https://3345423f-8bbf-4215-8004-215a3da67b60.users.dploy.ai/"

def resize_img(img, BASEWIDTH):
    wpercent = (BASEWIDTH / float(img.shape[0]))
    hsize = int(img.shape[1] * wpercent)
    wsize = int(img.shape[0] * wpercent)
    dim = (hsize, wsize)
    img = cv2.resize(img, dim, interpolation = cv2.INTER_AREA)
    return img


if __name__ == '__main__':
    # set up our tokens
    # api_key = os.environ['apikey']
    # user_id = os.environ['userid']
    api_key = "L4XSWSKVMFBTBFSLMTWBZ7I6LQ"
    user_id = "auth0|5ecb72629c5cbe0c1ac5c11e"

    # parse our command-line argument
    file_name = (sys.argv[-1])
    extension = file_name.split(".")[1]

    with open(file_name, "rb") as img:
        encoded_image = base64.b64encode(img.read())

    # set headers to include tokens (get them at api.dploy.ai)
    headers = {
            'content-type': 'application/json',
            'x-api-key': api_key,
            'x-api-user': user_id,
            }
    r = requests.post(url = dploy_url, json={"image": encoded_image, "type": extension}, headers = headers)
    data = json.loads(r.text)


    # get bounding boxes and results
    bbox, result, annotated_img = data['detected_face_coordinates'], data['detected_face_labels'], data['annotated_image']
    not_masked = list(map(lambda k: k[0], filter(lambda x: x[1] == 'not masked', (zip(bbox, result)))))
    print(not_masked)

    imdata = base64.b64decode(str(annotated_img))
    returned_img = Image.open(io.BytesIO(imdata))

    open_cv_img = cv2.imread(file_name)
    if open_cv_img.shape[0] != returned_img.width:
        open_cv_img = resize_img(open_cv_img, returned_img.width)

    print(open_cv_img.shape)
        


    kernel = (41,41) # how blurry it will be
    for face in not_masked: 
        x1, y1, x2, y2 = int(face[0]),int(face[1]),int(face[2]),int(face[3])
        # rescale it 
        open_cv_img[y1:y2, x1:x2] = cv2.blur(open_cv_img[y1:y2, x1:x2], kernel)

    cv2.imwrite("./blur.jpeg", open_cv_img)
    cv2.imshow("resized", open_cv_img)
    cv2.waitKey(0)


    
