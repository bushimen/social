import os
import re
import sys
import time
from datetime import datetime, timezone

import requests
from instagram_web_api import Client


class InstagramCrawler:
    api = None

    user_id = '327416611'

    social_endpoint = None

    def __init__(self):
        self.api = Client(auto_patch=True, drop_incompat_keys=False)
        self.social_endpoint = os.getenv('SOCIAL_ENDPOINT', 'http://localhost:8080')

    def fetch(self, end_cursor=None):
        result = self.api.user_feed(
            self.user_id,
            count=50,
            extract=False,
            end_cursor=end_cursor
        )
        
        info = self.parse_http_result(result)

        for post in info['posts']:
            self.process_post(post)
        
        page_info = info['page_info']
        if page_info.get('has_next_page', False):
            time.sleep(2)
            self.fetch(page_info['end_cursor'])

    
    def parse_http_result(self, result):
        status = result.get('status', 'error')

        if status != 'ok':
            sys.exit('api response not ok')

        data = result['data']
        media = data['user']['edge_owner_to_timeline_media']

        return {
            'count': media['count'],
            'posts': [edge['node'] for edge in media['edges']],
            'page_info': media['page_info']
        }

    def process_post(self, post):
        payload = self.parse_post(post)

        r = requests.put('%s/instagram' % self.social_endpoint, json = payload)

        if r.status_code != 200:
            sys.exit(r.text)
    
    def parse_post(self, post):
        text_edges = post['edge_media_to_caption']['edges']

        if not text_edges:
            caption = ''
            tags = []
        else:
            text = text_edges[0]['node']['text']

            text = re.sub('\s+', ' ', text)
            text = re.sub('\.\s+', '', text)
            tags = list({tag.strip().lower() for tag in re.findall('(?<=#)[^# ]+(?=#|$| )', text)})
            caption = re.sub('(#[^# ]+ )*(#[^# ]+$)', '', text)

        return {
            'shortcode': post['shortcode'],
            'caption': caption,
            'tags': tags,
            'likes': post['likes']['count'],
            'comments': post['comments']['count'],
            'type': post['type'],
            'thumbnail': post['images']['thumbnail']['url'],
            'image': post['images']['standard_resolution']['url'],
            'timestamp': datetime.utcfromtimestamp(int(post['created_time'])).replace(tzinfo=timezone.utc).isoformat()
        }


if __name__ == '__main__':
    crawler = InstagramCrawler()
    crawler.fetch()
