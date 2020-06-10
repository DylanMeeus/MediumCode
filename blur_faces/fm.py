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

def resize_img(img, new_width):
    wpercent = (new_width / float(img.shape[1]))
    wsize = int(img.shape[1] * wpercent)
    hsize = int(img.shape[0] * wpercent)
    dim = (wsize, hsize)
    img = cv2.resize(img, dim, interpolation = cv2.INTER_AREA)
    return img

def detect_unmasked(file_name):
    api_key = os.environ['apikey']
    user_id = os.environ['userid']
    with open(file_name, "rb") as img:
        encoded_image = base64.b64encode(img.read())

    # set headers to include tokens (get them at api.dploy.ai)
    headers = {
            'content-type': 'application/json',
            'x-api-key': api_key,
            'x-api-user': user_id,
            }
    r = requests.post(url = dploy_url, json={"image": encoded_image, "type": extension}, headers = headers)
    print(r.text)
    data = json.loads(r.text)


    # get bounding boxes and results
    bbox, result, annotated_img = data['detected_face_coordinates'], data['detected_face_labels'], data['annotated_image']
    not_masked_bbox = list(map(lambda k: k[0], filter(lambda x: x[1] == 'not masked', (zip(bbox, result)))))
    return (not_masked_bbox, annotated_img)


if __name__ == '__main__':
    # parse our command-line argument
    file_name = (sys.argv[-1])
    extension = file_name.split(".")[1]

    not_masked, annotated_img = detect_unmasked(file_name)

    imdata = base64.b64decode(str(annotated_img))
    returned_img = Image.open(io.BytesIO(imdata))

    open_cv_img = cv2.imread(file_name)
    if open_cv_img.shape[0] != returned_img.width:
        open_cv_img = resize_img(open_cv_img, returned_img.width)


    kernel = (41,41) # how blurry it will be
    for face in not_masked: 
        x1, y1, x2, y2 = int(face[0]),int(face[1]),int(face[2]),int(face[3])
        # rescale it 
        open_cv_img[y1:y2, x1:x2] = cv2.blur(open_cv_img[y1:y2, x1:x2], kernel)

    cv2.imwrite(file_name + "_blurred" + "." + extension, open_cv_img)
    cv2.imshow("resized", open_cv_img)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


    
