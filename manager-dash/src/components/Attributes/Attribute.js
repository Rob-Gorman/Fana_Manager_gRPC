import { useParams } from 'react-router-dom';
import { useEffect, useState, useCallback } from 'react';
import apiClient from '../../lib/apiClient';
import { attrTypeMapper } from '../../lib/formConstants';
import { AttributeAudience } from './AttributeAudience';
import DeleteIcon from '@mui/icons-material/Delete';

import { Box, Button, Stack, List, Divider, Typography, Grid } from '@mui/material';

const testAttribute = {
  attribute: 'beta',
  type: 'BOOL',
  createdAt: new Date(),
  audiences: [
    {
      displayName: 'Beta Testers',
      key: 'beta-testers',
      id: 1,
    },
    {
      displayName: 'Millenial Men',
      key: 'millenial-men',
      id: 2,
    },
  ],
};

export const Attribute = () => {
  const attrId = useParams().id;
  const [ready, setReady] = useState(false);
  const [attribute, setAttribute] = useState(testAttribute);

  useEffect(() => setReady(true));
  
  // const fetchAttribute = useCallback(async () => {
  //   const attr = await apiClient.getAttribute(attrId);
  //   setAttribute(attr);
  //   return attr;
  // }, [attrId])
  // useEffect(() => {
  //   const initialize = async () => {
  //     const attr = await fetchAttribute();
  //     setReady(true);
  //   }
  //   initialize();
  // }, [fetchAttribute])

  const handleDelete = async () => {
    try {
      const response = await apiClient.deleteAttribute(attrId);
    } catch (err) {
      console.error();
    }
  };

  if (!ready) {
    return <>Loading...</>;
  }

  return (
    <Box
      container="true"
      spacing={2}
      sx={{
        marginLeft: 8,
        maxWidth: 1000,
      }}
    >
      <Stack container="true" spacing={3}>
        <Typography variant="h3">Attribute Details</Typography>

     <Grid container>
      <Grid item xs={10}>


            <Stack>
              <Typography variant="caption">Title</Typography>
              <Typography variant="subtitle1">{attribute.attribute}</Typography>
            </Stack>
            <Stack>
              <Typography variant="caption">Type</Typography>
              <Typography variant="subtitle1">{attrTypeMapper[attribute.type]}</Typography>
            </Stack>
      </Grid>
   <Grid item xs={2} direction="column" alignItems="flex-end" justify="flex-end">

          <Button
            variant="outlined"
            onClick={handleDelete}
            startIcon={<DeleteIcon />}
            color="secondary"
            >
            Delete attribute
          </Button>
            </Grid>
     </Grid>

        <Stack container="true">
          <Typography variant="h4">Related Audiences</Typography>
          <Typography variant="subtitle2">
            List of audiences that reference this attribute
          </Typography>
        </Stack>
        <Stack
          container="true"
          divider={<Divider orientation="vertical" flexItem />}
          spacing={10}
          direction="row"
        >
          <Stack>
            <List style={{ width: 350 }}>
              {attribute.audiences.map((aud) => {
                return <AttributeAudience audience={aud} />;
              })}
            </List>
          </Stack>
        </Stack>
      </Stack>
    </Box>
  );
};
