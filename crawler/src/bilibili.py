import os
import re
import sys
import time
from datetime import datetime, timezone

import requests


class BilibiliCrawler:
    user_id = '24764396'

    page_current = 1

    page_total = 1

    social_endpoint = None

    def __init__(self):
        self.social_endpoint = os.getenv('SOCIAL_ENDPOINT', 'http://localhost:8080')

    def fetch(self, end_cursor=None):
        while self.page_current <= self.page_total:
            videos = self.fetch_page()

            for video in videos:
                video = self.fetch_video(video['aid'])
                self.process_video(video)
                time.sleep(1)
            
            self.page_current += 1

    def fetch_page(self):
        print('Fetching page %d' % self.page_current)

        r = requests.get('https://space.bilibili.com/ajax/member/getSubmitVideos?mid=%s&pagesize=50&order=pubdate&page=%d' %
                         (self.user_id, self.page_current))

        if r.status_code != 200:
            sys.exit(r.text)

        res = r.json()

        self.page_total = res['data']['pages']

        return res['data']['vlist']

    def fetch_video(self, aid):
        print('Fetching video %s' % aid)
        r = requests.get(
            'https://api.bilibili.com/x/web-interface/view?aid=%s' % aid)

        if r.status_code != 200:
            sys.exit(r.text)

        json = r.json()

        video = json['data']
        stat = video['stat']

        return {
            'aid': aid,
            'title': video['title'],
            'description': video['desc'],
            'favourites': stat['favorite'],
            'coins': stat['coin'],
            'likes': stat['like'],
            'dislikes': stat['dislike'],
            'danmaku': stat['danmaku'],
            'comments': stat['reply'],
            'shares': stat['share'],
            'views': stat['view'],
            'duration': video['duration'],
            'thumbnail': video['pic'] + '@380w_240h_100Q_1c.jpg',
            'image': video['pic'],
            'timestamp': datetime.utcfromtimestamp(video['pubdate']).replace(tzinfo=timezone.utc).isoformat()
        }

    def process_video(self, video):
        print('Processing video %s' % video['aid'])
        r = requests.put('%s/bilibili' % self.social_endpoint, json=video)

        if r.status_code != 200:
            sys.exit(r.text)


if __name__ == '__main__':
    crawler = BilibiliCrawler()
    crawler.fetch()
