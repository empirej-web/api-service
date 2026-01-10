import logging
import os
import json
from datetime import datetime, timedelta
from typing import Dict, List

class Utils:
    def __init__(self):
        self.logger = logging.getLogger(__name__)

    def load_config(self, file_path: str) -> Dict:
        try:
            with open(file_path, 'r') as file:
                return json.load(file)
        except FileNotFoundError:
            self.logger.error(f"Config file not found: {file_path}")
            return {}
        except json.JSONDecodeError as e:
            self.logger.error(f"Failed to parse config file: {e}")
            return {}

    def write_config(self, file_path: str, config: Dict) -> None:
        try:
            with open(file_path, 'w') as file:
                json.dump(config, file, indent=4)
        except Exception as e:
            self.logger.error(f"Failed to write config file: {e}")

    def get_current_time(self) -> str:
        return datetime.now().strftime("%Y-%m-%d %H:%M:%S")

    def get_time_diff(self, start_time: str, end_time: str) -> timedelta:
        start = datetime.strptime(start_time, "%Y-%m-%d %H:%M:%S")
        end = datetime.strptime(end_time, "%Y-%m-%d %H:%M:%S")
        return end - start

    def is_file_exists(self, file_path: str) -> bool:
        return os.path.isfile(file_path)

    def get_file_size(self, file_path: str) -> int:
        if self.is_file_exists(file_path):
            return os.path.getsize(file_path)
        else:
            return 0

    def get_file_extension(self, file_path: str) -> str:
        return os.path.splitext(file_path)[1]

def main():
    utils = Utils()
    config = utils.load_config('config.json')
    print(config)

if __name__ == "__main__":
    main()