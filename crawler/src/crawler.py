#seed url
#download the html of tha page that this url leads to
#extract all urls that are present on that html page
#store if they are not already craweled and repeate steps until no new urls are found
import logging
from urllib.parse import urljoin
import requests
from bs4 import BeautifulSoup
from collections import deque
import threading

logging.basicConfig(
    format='%(asctime)s %(levelname)-8s %(message)s',
    level=logging.INFO,
    datefmt='%Y-%m-%d %H:%M:%S'
)

class Crawler:
    def __init__(self, seed_url):
        self.seed_url = seed_url
        self.url_queue = deque()
        self.visited_urls = []
        for url in seed_url:
            self.url_queue.append(url)
    
    #must download the content on the page in order to get the links from it
    def download_url(self, url): 
        return requests.get(url).text
    
    def get_linked_urls(self, url, html):
        soup = BeautifulSoup(html, 'html.parser')
        for link in soup.find_all('a'):
            path = link.get('href')
            if path and path.startswith('/'):
                path = urljoin(url, path)
            yield path
    
    def add_url_to_visit(self, url):
        if url not in self.visited_urls and url not in self.url_queue:
            return url

    def crawl(self, url_to_crawl):
        html = self.download_url(url_to_crawl)
        url_list = []
        for url in self.get_linked_urls(url_to_crawl, html):
            curr_url = self.add_url_to_visit(url)
            if(curr_url):
                url_list.append(curr_url)
        
        return url_list


    def crawl_all_urls(self):
        while self.url_queue:
            url_to_crawl = self.url_queue.popleft()
            print(url_to_crawl)
            links_to_crawl = self.crawl(url_to_crawl)
            for link in links_to_crawl:
                self.url_queue.append(link)
                self.visited_urls.append(url_to_crawl)



if __name__ == '__main__':
    Crawler(seed_url=['https://go101.org/article/channel-closing.html']).crawl_all_urls()