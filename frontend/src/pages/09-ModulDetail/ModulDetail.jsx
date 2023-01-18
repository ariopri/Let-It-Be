import {
  Box,
  Container,
  Flex,
  Heading,
  List,
  ListItem,
  Stack,
  Text,
  useColorModeValue,
} from '@chakra-ui/react';
import { VscPlayCircle } from 'react-icons/vsc';
import { useParams } from 'react-router-dom';
import YouTube from 'react-youtube';
import axios from 'axios';
import { useState, useEffect } from 'react';
import Loading from '../../components/02-Reusable/LoadingEffect/LoadingFetchEffect';
import { ENDPOINT_API_GET_MODUL_DETAIL } from '../../api/api';

export default function ModulDetail() {
  const { name, id } = useParams();
  const [videoList, setVideoList] = useState([]);
  const [currentVideo, setCurrentVideo] = useState(null);
  const [description, setDescription] = useState('');
  const isDesktop = window.innerWidth >= 992;
  const isTablet = window.innerWidth >= 768 && window.innerWidth < 992;
  const isMobile = window.innerWidth < 768;
  const color = useColorModeValue('white', 'dark');
  const color2 = useColorModeValue('white', 'dark');
  const opts = {
    playerVars: {
      autoplay: 1,
    },
  };

  if (isDesktop) {
    opts.width = '640';
    opts.height = '390';
  } else if (isTablet) {
    opts.width = '480';
    opts.height = '270';
  } else if (isMobile) {
    opts.width = '380';
    opts.height = '200';
  }

  useEffect(() => {
    axios
      .get(ENDPOINT_API_GET_MODUL_DETAIL)
      .then((response) => {
        const data = response.data;
        const currentModul = data[0].pelajaran.find(
          (pelajaran) => pelajaran.name === name
        );
        const currentKelas = currentModul.kelas.find(
          (kelas) => kelas.id === Number(id)
        );
        setVideoList(currentKelas.video);
        setCurrentVideo(currentKelas.video[0]);
        setDescription(currentKelas.description);
      })
      .catch((error) => {
        console.log(error);
      });
  }, [id, name]);

  if (!currentVideo) {
    return <Loading />;
  }

  return (
    <Container maxW={'7xl'}>
      <Stack
        align={'center'}
        spacing={{ base: 8, md: 10 }}
        py={{ base: 20, md: 28 }}
        direction={{ base: 'column', md: 'row' }}
      >
        <Flex
          flex={2}
          justify="center"
          align="center"
          position="relative"
          w="full"
          flexWrap="wrap"
        >
          {currentVideo && (
            <YouTube videoId={currentVideo.videoid} opts={opts} />
          )}
        </Flex>

        <Stack flex={1} spacing={{ base: 5, md: 10 }}>
          <Heading fontSize={{ base: 'xl', md: '2xl', lg: '3xl' }}>
            {description}
          </Heading>

          <Stack
            spacing={{ base: 4, sm: 6 }}
            direction={{ base: 'column', sm: 'row' }}
          >
            <List styleType="none" m={0} p={0}>
              {videoList.map((video, index) => {
                return (
                  <ListItem key={index}>
                    <Flex
                      align="center"
                      p={2}
                      cursor="pointer"
                      onClick={() => setCurrentVideo(video)}
                    >
                      <Box mr={2}>
                        {currentVideo === video ? (
                          <VscPlayCircle name="play" size="24px" />
                        ) : (
                          <VscPlayCircle
                            name="play"
                            size="24px"
                            opacity="0.3"
                          />
                        )}
                      </Box>
                      <Box mr={2}>
                        <img
                          src={video.thumbnail}
                          alt={video.judul}
                          width="50px"
                          height="50px"
                        />
                      </Box>
                      <Text
                        color={currentVideo === video ? { color } : { color2 }}
                        fontWeight={currentVideo === video ? 'bold' : 'normal'}
                      >
                        {video.judul}
                      </Text>
                    </Flex>
                  </ListItem>
                );
              })}
            </List>
          </Stack>
        </Stack>
      </Stack>
    </Container>
  );
}
